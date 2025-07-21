// Package main implements the entry point for the swutrack application.
// It initializes and starts the HTTP server using the httpapi transport layer
// with support for graceful shutdown on OS signals (SIGINT, SIGTERM).
package main

import (
	"context"
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

// main initializes and starts the HTTP server on port 8080 with graceful shutdown support.
// It handles OS signals (SIGINT, SIGTERM) to ensure clean termination of all services.
func main() {
	// Define the server address
	const serverAddr = ":8080"

	// Log server startup information
	log.Printf("Starting swutrack application")
	log.Printf("Available endpoints:")
	log.Printf("  GET http://localhost:8080/hello")

	// Create the HTTP server
	server, err := httpapi.NewServer(serverAddr)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Set up channels for coordination
	shutdownChan := make(chan struct{})
	var wg sync.WaitGroup

	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create shutdown context with timeout
	const shutdownTimeout = 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// Create error channel with buffer size matching number of services
	services := []service{server}
	errChan := make(chan error, len(services))

	// Start all services
	startServices(ctx, shutdownChan, &wg, errChan, services...)

	// Wait for shutdown signal or fatal error
	select {
	case sig := <-sigChan:
		log.Printf("Received signal: %v, initiating graceful shutdown...", sig)
	case err := <-errChan:
		log.Printf("Fatal error: %v, initiating shutdown...", err)
		// Drain any additional errors to prevent goroutine blocking
		go func() {
			for err := range errChan {
				log.Printf("Additional error during shutdown: %v", err)
			}
		}()
	}

	// Trigger shutdown for all services
	close(shutdownChan)

	// Wait for all services to stop
	wg.Wait()

	// Close error channel to signal the drain goroutine to exit
	close(errChan)

	log.Printf("All services stopped, exiting")
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
