package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"swutrack/httpapi"
	"swutrack/svcerr"
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

func main() {
	log.Println("starting up...")
	ctx := context.Background()
	sctx, shutdown := context.WithCancel(ctx)
	errChan := make(chan svcerr.ServiceErr)
	var wg sync.WaitGroup
	conf := loadConfig()
	sigChan := createSigChan()

	apiService := httpapi.New(conf.Port)
	startServices(ctx, sctx, &wg, errChan, apiService)

	waitGracefulShutdown(sctx, shutdown, errChan, sigChan, &wg, conf)
}

// loadConfig loads and validates configuration from environment variables.
func loadConfig() (conf config) {
	if err := snest.Load(&conf); err != nil {
		log.Fatalf("failed to load environment variables: %s", err.Error())
	}

	if conf.Port == "" {
		log.Fatalf("Port cannot be empty")
	}

	return conf
}

// createSigChan creates and returns a buffered channel for OS signals.
func createSigChan() (sigChan chan os.Signal) {
	sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	return sigChan
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

func waitGracefulShutdown(sctx context.Context, shutdown context.CancelFunc, errChan chan svcerr.ServiceErr, sigChan chan os.Signal, wg *sync.WaitGroup, conf config) {
	select {
	case err := <-errChan:
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
