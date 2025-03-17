package httpapi

import (
	"io"
	"net/http"
	"testing"
)

func Test_Hello(t *testing.T) {
	// Setup expectations.
	expected := "Hello, world!"

	// Do.
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		t.Fatalf("failed to get endpoint: %s", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %s", err)
	}
	bodyString := string(body)

	// Assert results.
	if bodyString != "Hello, world!" {
		t.Fatalf("unexpected response\nexpected: %s\nactual: %s", expected, bodyString)
	}
}
