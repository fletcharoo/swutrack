// Package main implements a minimal HTTP server for the swutrack application.
// This server provides a basic "hello world" endpoint as a foundation for future API development.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler handles GET requests to the /hello endpoint and returns "hello world".
func helloHandler(w http.ResponseWriter, r *http.Request) (err error) {
	// Set the content type to plain text
	w.Header().Set("Content-Type", "text/plain")
	
	// Write the response
	_, err = fmt.Fprint(w, "hello world")
	if err != nil {
		err = fmt.Errorf("failed to write response: %w", err)
		return
	}
	
	return
}

// main initializes and starts the HTTP server on port 8080.
func main() {
	// Create a new multiplexer for routing
	mux := http.NewServeMux()
	
	// Register the hello handler for the /hello endpoint
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if err := helloHandler(w, r); err != nil {
			log.Printf("Error handling request: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
	
	// Define the server address
	const serverAddr = ":8080"
	
	// Log server startup information
	log.Printf("Starting HTTP server on port 8080")
	log.Printf("Available endpoints:")
	log.Printf("  GET http://localhost:8080/hello")
	
	// Create and start the HTTP server
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}
	
	// Start the server and handle any startup errors
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server on port 8080: %v", err)
	}
}