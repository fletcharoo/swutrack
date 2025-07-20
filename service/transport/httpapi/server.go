// Package httpapi provides the HTTP transport layer for the swutrack application.
// It handles HTTP server configuration, request routing, and handler management.
package httpapi

import (
	"fmt"
	"net/http"
)

// Server encapsulates the HTTP server and its configuration.
type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
}

// NewServer creates and initializes a new HTTP server instance.
// The addr parameter specifies the TCP address for the server to listen on.
func NewServer(addr string) (server *Server, err error) {
	if addr == "" {
		err = fmt.Errorf("server address cannot be empty")
		return
	}

	mux := http.NewServeMux()

	server = &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
		mux: mux,
	}

	// Set up all routes
	err = server.setupRoutes()
	if err != nil {
		err = fmt.Errorf("failed to setup routes: %w", err)
		return
	}

	return server, nil
}

// Start begins listening for HTTP requests on the configured address.
// This method blocks until the server is shut down or encounters an error.
func (s *Server) Start() (err error) {
	if s == nil {
		err = fmt.Errorf("server cannot be nil")
		return
	}

	if s.httpServer == nil {
		err = fmt.Errorf("HTTP server is not initialized")
		return
	}

	err = s.httpServer.ListenAndServe()
	if err != nil {
		err = fmt.Errorf("failed to start HTTP server: %w", err)
		return
	}

	return nil
}
