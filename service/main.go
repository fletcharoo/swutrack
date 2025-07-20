// Package main implements the entry point for the swutrack application.
// It initializes and starts the HTTP server using the httpapi transport layer.
package main

import (
	"log"
	"swutrack/service/transport/httpapi"
)

// main initializes and starts the HTTP server on port 8080.
func main() {
	// Define the server address
	const serverAddr = ":8080"

	// Log server startup information
	log.Printf("Starting HTTP server on port 8080")
	log.Printf("Available endpoints:")
	log.Printf("  GET http://localhost:8080/hello")

	// Create the HTTP server
	server, err := httpapi.NewServer(serverAddr)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Start the server and handle any startup errors
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server on port 8080: %v", err)
	}
}
