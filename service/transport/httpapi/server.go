// Package httpapi provides the HTTP transport layer for the swutrack application.
// It handles HTTP server configuration, request routing, and handler management.
package httpapi

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Server encapsulates the HTTP server and its configuration.
// It implements the service interface for lifecycle management.
type Server struct {
	httpServer *http.Server   // The underlying HTTP server instance
	mux        *http.ServeMux // HTTP request multiplexer for routing
	mu         sync.Mutex     // Protects concurrent access to isRunning
	isRunning  bool           // Tracks whether the server is currently running
}

// handlerFunc represents a handler function that can return an error.
type handlerFunc func(http.ResponseWriter, *http.Request) error

// serverConfig holds the configuration for creating a new server.
type serverConfig struct {
	addr         string
	helloHandler handlerFunc
}

// NewServer creates and initializes a new HTTP server instance.
// The addr parameter specifies the TCP address for the server to listen on.
func NewServer(addr string) (server *Server, err error) {
	if addr == "" {
		err = fmt.Errorf("server address cannot be empty")
		return
	}

	// Use the internal constructor with production handlers
	config := serverConfig{
		addr:         addr,
		helloHandler: helloHandler,
	}

	return newServerWithConfig(config)
}

// newServerWithConfig creates a new server with the provided configuration.
// This is an internal constructor used for testing with dependency injection.
func newServerWithConfig(config serverConfig) (server *Server, err error) {
	if config.addr == "" {
		err = fmt.Errorf("server address cannot be empty")
		return
	}

	mux := http.NewServeMux()

	server = &Server{
		httpServer: &http.Server{
			Addr:    config.addr,
			Handler: mux,
		},
		mux: mux,
	}

	// Set up all routes with the provided handlers
	err = server.setupRoutesWithHandlers(config.helloHandler)
	if err != nil {
		err = fmt.Errorf("failed to setup routes: %w", err)
		server = nil
		return
	}

	return server, nil
}

// Name returns the service's display name for logging purposes.
func (s *Server) Name() string {
	return "HTTP Server"
}

// Start begins the service operation, sending any errors to the provided channel.
// This method starts the server in a non-blocking way and returns immediately.
func (s *Server) Start(errChan chan error) {
	if s == nil {
		errChan <- fmt.Errorf("server cannot be nil")
		return
	}

	if s.httpServer == nil {
		errChan <- fmt.Errorf("HTTP server is not initialized")
		return
	}

	// Check if server is already running
	s.mu.Lock()
	if s.isRunning {
		s.mu.Unlock()
		errChan <- fmt.Errorf("server is already running")
		return
	}

	// Mark server as running
	s.isRunning = true
	s.mu.Unlock()

	// Start the server in a goroutine
	go func() {
		log.Printf("HTTP Server listening on %s", s.httpServer.Addr)
		err := s.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("HTTP server stopped unexpectedly: %w", err)
		}
		// Server has stopped
		s.mu.Lock()
		s.isRunning = false
		s.mu.Unlock()
	}()
}

// Stop gracefully shuts down the service within the provided context deadline.
func (s *Server) Stop(ctx context.Context) (err error) {
	if s == nil {
		err = fmt.Errorf("server cannot be nil")
		return
	}

	if s.httpServer == nil {
		err = fmt.Errorf("HTTP server is not initialized")
		return
	}

	log.Printf("HTTP Server shutting down...")
	err = s.httpServer.Shutdown(ctx)
	if err != nil {
		err = fmt.Errorf("failed to shutdown HTTP server: %w", err)
		return
	}

	log.Printf("HTTP Server shutdown complete")
	return nil
}
