// Package httpapi provides an HTTP server implementation for the swutrack
// service.
//
// httpapi exposes a Server type that implements a base HTTP API server
// with configurable port and graceful shutdown capabiilties.
// Server implements the service interface defined in main, allowing it to be
// managed alongside other services in the application.
package httpapi

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"swutrack/svcerr"
)

// Service represents an HTTP server instance that can be started and stopped.
type Service struct {
	port   string
	server *http.Server
}

// New creates and returns a new Service instance with the specified port.
// The port parameter should be a valid port number as a string (e.g. "8080").
// The returned Service will not start listening until Start is called.
func New(port string) *Service {
	return &Service{
		port: port,
	}
}

// Name returns the name of the service.
func (s Service) Name() (name string) {
	return "HTTP API"
}

// Start initializes and starts the HTTP server on the configured port.
// If the server encounters a non-graceful shutdown error, it will be sent to
// the provided error channel.
// The function blocks until the server is stopped or encounters an error.
func (s *Service) Start(errChan chan svcerr.ServiceErr) {
	if s == nil {
		errChan <- svcerr.New(s.Name(), fmt.Errorf("pointer cannot be nil"))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handleHello)

	s.server = &http.Server{
		Addr:    "0.0.0.0:" + s.port,
		Handler: mux,
	}

	log.Printf("serving %q on %s", s.Name(), s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		errChan <- svcerr.New(s.Name(), err)
		return
	}
}

// Stop gracefully shuts down the HTTP server.
// It waits for active connections to complete or until the provided context is
// cancelled.
// If the shutdown fails or the context is cancelled before completion, Stop
// returns an error.
func (s *Service) Stop(ctx context.Context) (err svcerr.ServiceErr) {
	return svcerr.New(s.Name(), s.server.Shutdown(ctx))
}

// handleHello responds to HTTP requests with a simple "Hello, world!" message.
// It writes the message directly to the provided http.ResponseWriter.
func handleHello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
