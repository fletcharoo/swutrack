// Package svcerr provides error handling functionality for services in the
// swutrack application.
//
// svcerr exposes a ServiceErr type that wraps errors with service name context
// allowing errors to be traced back to their originating service.
// The reason this package is used instead of a standard error type is to
// ensure the caller has context of which service the error came from.
//
// Example usage:
//
//	if err != nil {
//	    return svcerr.New("MyService", err)
//	}
package svcerr

import (
	"fmt"
	"log"
)

// ServiceErr represents a service error that includes the name of the service
// that generated the error along with the underlying error itself.
type ServiceErr struct {
	name string
	err  error
}

// Error returns a string representation of the underlying error in the form of
// "{service name}: {error}".
func (se ServiceErr) Error() string {
	return fmt.Sprintf("%s: %s", se.name, se.err)
}

// HasError returns whether the underlying error is non-nil.
func (se ServiceErr) HasError() bool {
	return se.err != nil
}

// New creates a new ServiceErr with the given service name and underlying
// error.
// If the name parameter is empty, it logs a warning.
func New(name string, err error) (svcerr ServiceErr) {
	if name == "" {
		log.Println("ServiceErr.name cannot be empty")
	}

	return ServiceErr{
		name: name,
		err:  err,
	}
}
