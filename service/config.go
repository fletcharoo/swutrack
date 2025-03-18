package main

import "time"

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
