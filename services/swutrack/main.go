package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"swutrack/services/swutrack/httpapi"
	"swutrack/services/swutrack/svcerr"
	"sync"
	"syscall"
	"time"

	"github.com/fletcharoo/snest"
)

// service defines the interface for manageable services in the applicatino.
type service interface {
	Name() string
	Start(chan svcerr.ServiceErr)
	Stop(context.Context) svcerr.ServiceErr
}

// config holds the application configuration loaded from environment
// variables.
type config struct {
	// Port specifies the network port that the HTTP server will listen on.
	// This field is required and must not be empty.
	Port string `snest:"PORT"`

	// ShutdownTimeout specifies how long to wait for services to gracefully
	// shutdown before forcefully terminating them.
	ShutdownTimeout time.Duration `snest:"SHUTDOWN_TIMEOUT"`
}

// loadConfig loads and validates configuration from environment variables.
func loadConfig() (conf config, err error) {
	if err = snest.Load(&conf); err != nil {
		err = fmt.Errorf("failed to load environment variables: %w", err)
		return
	}

	if conf.Port == "" {
		err = fmt.Errorf("Port cannot be empty")
		return
	}

	return conf, nil
}

func main() {
	// Setup.
	log.Println("starting up...")
	ctx := context.Background()
	sctx, shutdown := context.WithCancel(ctx)
	errChan := make(chan svcerr.ServiceErr)
	var wg sync.WaitGroup

	// Load configuration.
	conf, err := loadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	// Listen for OS signals.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Setup services.
	apiService := &httpapi.Service{
		Port: conf.Port,
	}

	startServices(ctx, sctx, &wg, errChan, apiService)

	// Graceful shutdown.
	select {
	case err = <-errChan:
		log.Printf("service error: %s", err.Error())
		shutdown()
	case <-sigChan:
		shutdown()
	case <-sctx.Done():
	}

	log.Println("shutting down...")
	go func() {
		select {
		case <-time.After(conf.ShutdownTimeout):
			log.Fatalf("shutdown timeout reached")
		}
	}()

	wg.Wait()
	log.Println("shutdown successful")
}

// startServices starts the provided services and manages their lifecycle.
// If a service fails to start, it pushes the error on the error channel.
// If a service fails to stop, it logs the error.
func startServices(ctx context.Context, sctx context.Context, wg *sync.WaitGroup, errChan chan svcerr.ServiceErr, services ...service) {
	for _, s := range services {
		// Start the service.
		wg.Add(1)
		serviceName := s.Name()
		log.Printf("starting %q\n", serviceName)
		go s.Start(errChan)

		// Shutdown goroutine.
		go func() {
			defer wg.Done()
			<-sctx.Done()
			if err := s.Stop(ctx); err.HasError() {
				log.Printf("failed to stop service: %s\n", err.Error())
			}
		}()
	}
}
