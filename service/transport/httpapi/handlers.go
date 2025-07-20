package httpapi

import (
	"fmt"
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

	return nil
}
