package main

import (
	"fmt"
	"time"
)

// config holds the application configuration loaded from environment
// variables.
type config struct {
	// Port specifies the network port that the HTTP server will listen on.
	// This field is required and must not be empty.
	Port string `envconfig:"PORT"`

	// ShutdownTimeout specifies how long to wait for services to gracefully
	// shutdown before forcefully terminating them.
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT"`

	// PostgresConnURL specifies the connection URL for the PostgreSQL database.
	PostgresConnURL string `envconfig:"POSTGRES_CONN_URL"`
}

// validate checks that the configuration is valid.
func (c config) validate() (err error) {
	if c.Port == "" {
		err = fmt.Errorf("port cannot be empty")
		return
	}

	if c.ShutdownTimeout == 0 {
		err = fmt.Errorf("shutdown timeout cannot be empty")
		return
	}

	if c.PostgresConnURL == "" {
		err = fmt.Errorf("postgres connection URL cannot be empty")
		return
	}

	return nil
}
