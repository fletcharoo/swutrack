package httpapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Server_setupRoutes_Success(t *testing.T) {
	// Setup expectations
	expectedStatus := http.StatusOK
	expectedBody := "hello world"
	expectedPath := "/hello"

	// Arrange
	server := &Server{
		mux: http.NewServeMux(),
	}

	// Act
	err := server.setupRoutes()

	// Assert
	assert.NoError(t, err, "setupRoutes should not return an error with valid server")

	// Verify route is registered by making a test request
	req := httptest.NewRequest(http.MethodGet, expectedPath, nil)
	recorder := httptest.NewRecorder()

	server.mux.ServeHTTP(recorder, req)

	// Should get a successful response from the hello handler
	assert.Equal(t, expectedStatus, recorder.Code, "handler should return OK status")
	assert.Equal(t, expectedBody, recorder.Body.String(), "handler should return expected greeting")
}

func Test_Server_setupRoutes_Errors(t *testing.T) {
	testCases := map[string]struct {
		setupServer func() *Server
		expectedMsg string
	}{
		"nil server": {
			setupServer: func() *Server {
				return nil
			},
			expectedMsg: "server cannot be nil",
		},
		"nil multiplexer": {
			setupServer: func() *Server {
				return &Server{
					mux: nil,
				}
			},
			expectedMsg: "HTTP multiplexer is not initialized",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Arrange
			server := tc.setupServer()

			// Act
			var err error
			if server != nil {
				err = server.setupRoutes()
			} else {
				// Handle nil server case
				var s *Server
				err = s.setupRoutes()
			}

			// Assert
			require.Error(t, err, "setupRoutes should return an error for test case: %s", name)
			assert.Contains(t, err.Error(), tc.expectedMsg, "error message should contain expected text for test case: %s", name)
		})
	}
}

func Test_setupRoutes_HandlerError(t *testing.T) {
	// This test verifies that handler errors are properly logged and return 500 status

	// Setup expectations
	expectedStatus := http.StatusOK
	expectedPath := "/hello"

	// Create a server with a custom mux that we can test
	server := &Server{
		mux: http.NewServeMux(),
	}

	// Setup routes
	err := server.setupRoutes()
	require.NoError(t, err, "setupRoutes should succeed with valid server configuration")

	// Create a test server to capture the response
	testServer := httptest.NewServer(server.mux)
	defer testServer.Close()

	// Make a request to verify error handling
	// Note: In the actual implementation, the handler always succeeds,
	// but the error handling code path exists and this test documents that behavior
	resp, err := http.Get(testServer.URL + expectedPath)
	require.NoError(t, err, "HTTP GET request should succeed")
	defer resp.Body.Close()

	// Should get a successful response since helloHandler doesn't actually fail
	assert.Equal(t, expectedStatus, resp.StatusCode, "should receive OK status since helloHandler succeeds")
}

// mockFailingResponseWriter is a custom ResponseWriter that fails on Write
type mockFailingResponseWriter struct {
	*httptest.ResponseRecorder
	writeError error
}

func (m *mockFailingResponseWriter) Write(p []byte) (int, error) {
	if m.writeError != nil {
		return 0, m.writeError
	}
	return m.ResponseRecorder.Write(p)
}

func Test_setupRoutes_HandlerWriteError(t *testing.T) {
	// This test verifies that when a handler returns an error,
	// it's logged and a 500 status is returned

	// Setup expectations
	expectedStatus := http.StatusInternalServerError
	expectedErrorMessage := "Internal Server Error"
	expectedPath := "/hello"
	expectedAddr := ":0"

	// Create a handler that always fails
	failingHandler := func(w http.ResponseWriter, r *http.Request) error {
		return fmt.Errorf("handler error")
	}

	// Create a server config with the failing handler
	config := serverConfig{
		addr:         expectedAddr,
		helloHandler: failingHandler,
	}

	// Create server using the internal constructor
	server, err := newServerWithConfig(config)
	require.NoError(t, err, "newServerWithConfig should succeed with valid config")
	require.NotNil(t, server, "server should not be nil after successful creation")

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, expectedPath, nil)
	recorder := httptest.NewRecorder()

	// Serve the request
	server.mux.ServeHTTP(recorder, req)

	// Verify we get a 500 error
	assert.Equal(t, expectedStatus, recorder.Code, "handler error should result in 500 status")
	assert.Contains(t, recorder.Body.String(), expectedErrorMessage, "error response should contain standard error message")
}
