package httpapi

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockResponseWriter is a mock implementation of http.ResponseWriter that can simulate write errors.
type mockResponseWriter struct {
	httptest.ResponseRecorder
	writeError error
}

// Write overrides the Write method to return an error if configured.
func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	if m.writeError != nil {
		return 0, m.writeError
	}
	return m.ResponseRecorder.Write(p)
}

func Test_helloHandler_Success(t *testing.T) {
	// Setup expectations
	expectedCode := http.StatusOK
	expectedBody := "hello world"
	expectedType := "text/plain"
	expectedPath := "/hello"

	// Arrange
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, expectedPath, nil)

	// Act
	err := helloHandler(w, r)

	// Assert
	assert.NoError(t, err, "helloHandler should not return an error for successful response")
	assert.Equal(t, expectedCode, w.Code, "response code should match expected value")
	assert.Equal(t, expectedBody, w.Body.String(), "response body should contain expected content")
	assert.Equal(t, expectedType, w.Header().Get("Content-Type"), "Content-Type header should be set correctly")
}

func Test_helloHandler_WriteError(t *testing.T) {
	// Setup expectations
	expectedErrorMessage := "failed to write response"
	expectedPath := "/hello"

	// Arrange
	mock := &mockResponseWriter{
		ResponseRecorder: *httptest.NewRecorder(),
		writeError:       errors.New("write failed"),
	}
	r := httptest.NewRequest(http.MethodGet, expectedPath, nil)

	// Act
	err := helloHandler(mock, r)

	// Assert
	require.Error(t, err, "helloHandler should return an error when write fails")
	assert.Contains(t, err.Error(), expectedErrorMessage, "error message should indicate write failure")
}
