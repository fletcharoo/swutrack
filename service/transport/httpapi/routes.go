package httpapi

import (
	"fmt"
	"log"
	"net/http"
)

// setupRoutes configures all HTTP routes for the server.
// This method registers all endpoints and their corresponding handlers.
func (s *Server) setupRoutes() (err error) {
	if s == nil {
		err = fmt.Errorf("server cannot be nil")
		return
	}

	if s.mux == nil {
		err = fmt.Errorf("HTTP multiplexer is not initialized")
		return
	}

	// Register the hello handler for the /hello endpoint
	s.mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if err := helloHandler(w, r); err != nil {
			log.Printf("Error handling request: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	return nil
}
