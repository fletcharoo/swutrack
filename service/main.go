// Package main implements the entry point for the swutrack application.
// It initializes and starts the HTTP server using the httpapi transport layer
// with support for graceful shutdown on OS signals (SIGINT, SIGTERM).
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"swutrack/service/transport/httpapi"
	"sync"
	"syscall"
	"time"
)

// service defines the interface for manageable services in the application.
// Each service must provide its name, start method, and stop method for lifecycle management.
type service interface {
	// Name returns the service's display name for logging purposes.
	Name() string
	// Start begins the service operation, sending any errors to the provided channel.
	Start(chan error)
	// Stop gracefully shuts down the service within the provided context deadline.
	Stop(context.Context) error
}

// signalNotifier is an interface for signal notification (for testing).
type signalNotifier interface {
	// Notify registers the channel to receive notifications.
	Notify(c chan<- os.Signal, sig ...os.Signal)
	// Stop stops the notification.
	Stop(c chan<- os.Signal)
}

// osSignalNotifier is the production implementation of signalNotifier.
type osSignalNotifier struct{}

// Notify registers the channel to receive OS signal notifications.
func (osSignalNotifier) Notify(c chan<- os.Signal, sig ...os.Signal) {
	signal.Notify(c, sig...)
}

// Stop stops the OS signal notifications for the given channel.
func (osSignalNotifier) Stop(c chan<- os.Signal) {
	signal.Stop(c)
}

// serverFactory is a function type for creating servers (for testing).
type serverFactory func(addr string) (service, error)

// config holds the application configuration.
type config struct {
	serverAddr      string
	shutdownTimeout time.Duration
	serverFactory   serverFactory
	signalNotifier  signalNotifier
}

// defaultConfig returns the default production configuration.
func defaultConfig() config {
	return config{
		serverAddr:      ":8080",
		shutdownTimeout: 30 * time.Second,
		serverFactory: func(addr string) (service, error) {
			return httpapi.NewServer(addr)
		},
		signalNotifier: osSignalNotifier{},
	}
}

// main initializes and starts the HTTP server on port 8080 with graceful shutdown support.
// It handles OS signals (SIGINT, SIGTERM) to ensure clean termination of all services.
func main() {
	if err := run(defaultConfig()); err != nil {
		log.Fatalf("Application failed: %v", err)
	}
}

// run contains the main application logic with injected dependencies.
func run(cfg config) (err error) {
	// Log server startup information
	log.Printf("Starting swutrack application")
	log.Printf("Available endpoints:")
	log.Printf("  GET http://localhost%s/hello", cfg.serverAddr)

	// Create the HTTP server
	server, err := cfg.serverFactory(cfg.serverAddr)
	if err != nil {
		err = fmt.Errorf("failed to create server: %w", err)
		return
	}

	// Set up channels for coordination
	shutdownChan := make(chan struct{})
	var wg sync.WaitGroup

	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	cfg.signalNotifier.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	defer cfg.signalNotifier.Stop(sigChan)

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), cfg.shutdownTimeout)
	defer cancel()

	// Create error channel with buffer size matching number of services
	services := []service{server}
	errChan := make(chan error, len(services))

	// Start all services
	startServices(ctx, shutdownChan, &wg, errChan, services...)

	// Wait for shutdown signal or fatal error
	shutdownReason := waitForShutdown(sigChan, errChan)
	log.Printf("Shutdown reason: %s", shutdownReason)

	// Trigger shutdown for all services
	close(shutdownChan)

	// Wait for all services to stop
	wg.Wait()

	// Close error channel to signal any drain goroutine to exit
	close(errChan)

	log.Printf("All services stopped, exiting")
	return nil
}

// waitForShutdown waits for either a signal or an error and returns the reason.
func waitForShutdown(sigChan <-chan os.Signal, errChan <-chan error) (reason string) {
	select {
	case sig := <-sigChan:
		reason = fmt.Sprintf("received signal: %v", sig)
	case err := <-errChan:
		reason = fmt.Sprintf("fatal error: %v", err)
		// Drain any additional errors to prevent goroutine blocking
		go func() {
			for err := range errChan {
				log.Printf("Additional error during shutdown: %v", err)
			}
		}()
	}
	return reason
}

// startServices manages the lifecycle of multiple services concurrently.
// It starts each service in a separate goroutine and handles graceful shutdown
// when the shutdownChan is closed. Any errors during shutdown are logged.
func startServices(ctx context.Context, shutdownChan chan struct{}, wg *sync.WaitGroup, errChan chan error, services ...service) {
	for _, s := range services {
		wg.Add(1)
		serviceName := s.Name()
		log.Printf("starting: %q", serviceName)
		go s.Start(errChan)

		go func() {
			defer wg.Done()
			<-shutdownChan
			if err := s.Stop(ctx); err != nil {
				log.Printf("failed to stop service: %s: %s\n", serviceName, err.Error())
			}
		}()
	}
}
