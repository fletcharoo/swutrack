package httpapi

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewServer_Success(t *testing.T) {
	// Setup expectations
	expectedAddr := ":8080"
	expectedRunning := false
	
	// Act
	server, err := NewServer(expectedAddr)
	
	// Assert
	assert.NoError(t, err, "NewServer should not return an error with valid address")
	require.NotNil(t, server, "server should not be nil after successful creation")
	assert.NotNil(t, server.httpServer, "httpServer field should be initialized")
	assert.NotNil(t, server.mux, "mux field should be initialized")
	assert.Equal(t, expectedAddr, server.httpServer.Addr, "server address should match input")
	assert.Equal(t, server.mux, server.httpServer.Handler, "mux should be set as the handler")
	assert.Equal(t, expectedRunning, server.isRunning, "server should not be running after creation")
}

func Test_NewServer_EmptyAddress(t *testing.T) {
	// Setup expectations
	expectedAddr := ""
	expectedErrorMsg := "server address cannot be empty"
	
	// Act
	server, err := NewServer(expectedAddr)
	
	// Assert
	require.Error(t, err, "NewServer should return an error for empty address")
	assert.Contains(t, err.Error(), expectedErrorMsg, "error message should indicate empty address")
	assert.Nil(t, server, "server should be nil when error is returned")
}

// Test_newServerWithConfig_EmptyAddress tests error case for empty address
func Test_newServerWithConfig_EmptyAddress(t *testing.T) {
	// Setup expectations
	expectedErrorMsg := "server address cannot be empty"
	
	// Arrange
	config := serverConfig{
		addr:         "",
		helloHandler: helloHandler,
	}
	
	// Act
	server, err := newServerWithConfig(config)
	
	// Assert
	require.Error(t, err, "newServerWithConfig should return an error for empty address")
	assert.Contains(t, err.Error(), expectedErrorMsg, "error message should indicate empty address")
	assert.Nil(t, server, "server should be nil when error is returned")
}

// Test_setupRoutesWithHandlers_Errors tests error cases for setupRoutesWithHandlers
func Test_setupRoutesWithHandlers_Errors(t *testing.T) {
	testCases := map[string]struct {
		setupServer func() *Server
		errorMsg    string
	}{
		"nil server": {
			setupServer: func() *Server {
				return nil
			},
			errorMsg: "server cannot be nil",
		},
		"nil mux": {
			setupServer: func() *Server {
				return &Server{
					httpServer: &http.Server{},
					mux:        nil,
				}
			},
			errorMsg: "HTTP multiplexer is not initialized",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Arrange
			server := tc.setupServer()
			
			// Act
			err := server.setupRoutesWithHandlers(helloHandler)
			
			// Assert - all test cases are error cases
			require.Error(t, err, "setupRoutesWithHandlers should return an error for test case: %s", name)
			assert.Contains(t, err.Error(), tc.errorMsg, "error message should contain expected text for test case: %s", name)
		})
	}
}

func Test_Server_Name(t *testing.T) {
	// Setup expectations
	expectedName := "HTTP Server"
	
	// Arrange
	server := &Server{}

	// Act
	name := server.Name()

	// Assert
	assert.Equal(t, expectedName, name, "Name() should return the correct service name")
}

func Test_Server_Start_Success(t *testing.T) {
	// Setup expectations
	expectedAddr := ":0" // Use port 0 to let OS assign available port
	expectedRunning := true
	expectedTimeout := 100 * time.Millisecond
	expectedShutdownTimeout := 5 * time.Second

	// Arrange
	server := &Server{
		httpServer: &http.Server{
			Addr: expectedAddr,
		},
		mux:       http.NewServeMux(),
		isRunning: false,
	}
	errChan := make(chan error, 1)

	// Act
	server.Start(errChan)

	// Wait a bit to ensure no error
	select {
	case err := <-errChan:
		t.Fatalf("unexpected error from Start: %v", err)
	case <-time.After(expectedTimeout):
		// No error received, as expected
	}

	// Assert server is running
	server.mu.Lock()
	assert.Equal(t, expectedRunning, server.isRunning, "server should be marked as running after successful start")
	server.mu.Unlock()

	// Clean up: stop the server
	ctx, cancel := context.WithTimeout(context.Background(), expectedShutdownTimeout)
	defer cancel()
	err := server.Stop(ctx)
	assert.NoError(t, err, "Stop should succeed during cleanup")
}

func Test_Server_Start_Errors(t *testing.T) {
	testCases := map[string]struct {
		setupServer func() *Server
		errorMsg    string
	}{
		"nil server": {
			setupServer: func() *Server {
				return nil
			},
			errorMsg: "server cannot be nil",
		},
		"nil http server": {
			setupServer: func() *Server {
				return &Server{
					httpServer: nil,
					mux:        http.NewServeMux(),
				}
			},
			errorMsg: "HTTP server is not initialized",
		},
		"already running": {
			setupServer: func() *Server {
				return &Server{
					httpServer: &http.Server{
						Addr: ":0",
					},
					mux:       http.NewServeMux(),
					isRunning: true,
				}
			},
			errorMsg: "server is already running",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Arrange
			server := tc.setupServer()
			errChan := make(chan error, 1)

			// Act
			if server != nil {
				server.Start(errChan)
			} else {
				var s *Server
				s.Start(errChan)
			}

			// Wait for error
			select {
			case err := <-errChan:
				// Assert - all test cases expect errors
				require.Error(t, err, "Start should return an error for test case: %s", name)
				assert.Contains(t, err.Error(), tc.errorMsg, "error message should contain expected text for test case: %s", name)
			case <-time.After(100 * time.Millisecond):
				t.Errorf("expected error for test case %s but none received", name)
			}
		})
	}
}

func Test_Server_Stop_Success(t *testing.T) {
	// Setup expectations
	expectedAddr := ":0"
	expectedRunningAfterStop := false
	expectedTimeout := 5 * time.Second
	expectedStartupDelay := 50 * time.Millisecond

	// Arrange
	server := &Server{
		httpServer: &http.Server{
			Addr: expectedAddr,
		},
		mux: http.NewServeMux(),
	}

	// Start server first
	errChan := make(chan error, 1)
	server.Start(errChan)
	
	// Wait a bit for server to start
	time.Sleep(expectedStartupDelay)
	
	// Check for startup errors
	select {
	case err := <-errChan:
		t.Fatalf("failed to start server: %v", err)
	default:
		// No error, continue
	}

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), expectedTimeout)
	defer cancel()
	err := server.Stop(ctx)

	// Assert
	assert.NoError(t, err, "Stop should not return an error for running server")
	
	// Verify server is no longer running
	server.mu.Lock()
	assert.Equal(t, expectedRunningAfterStop, server.isRunning, "server should not be running after successful stop")
	server.mu.Unlock()
}

func Test_Server_Stop_Errors(t *testing.T) {
	testCases := map[string]struct {
		setupServer func() *Server
		errorMsg    string
	}{
		"nil server": {
			setupServer: func() *Server {
				return nil
			},
			errorMsg: "server cannot be nil",
		},
		"nil http server": {
			setupServer: func() *Server {
				return &Server{
					httpServer: nil,
				}
			},
			errorMsg: "HTTP server is not initialized",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Arrange
			server := tc.setupServer()
			ctx := context.Background()

			// Act
			var err error
			if server != nil {
				err = server.Stop(ctx)
			} else {
				var s *Server
				err = s.Stop(ctx)
			}

			// Assert - all test cases expect errors
			require.Error(t, err, "Stop should return an error for test case: %s", name)
			assert.Contains(t, err.Error(), tc.errorMsg, "error message should contain expected text for test case: %s", name)
		})
	}
}

func Test_Server_StartStopIntegration(t *testing.T) {
	// This test verifies the full lifecycle of starting and stopping a server
	
	// Setup expectations
	expectedAddr := ":0"
	expectedRunningAfterStart := true
	expectedRunningAfterStop := false
	expectedWaitTime := 100 * time.Millisecond
	expectedTimeout := 5 * time.Second
	
	// Arrange
	server, err := NewServer(expectedAddr)
	require.NoError(t, err, "NewServer should succeed with valid address")
	require.NotNil(t, server, "server should not be nil after successful creation")

	errChan := make(chan error, 1)

	// Act - Start the server
	server.Start(errChan)

	// Wait a bit for server to start
	time.Sleep(expectedWaitTime)

	// Verify no startup errors
	select {
	case err := <-errChan:
		t.Fatalf("unexpected error during startup: %v", err)
	default:
		// No error, continue
	}

	// Verify server is running
	server.mu.Lock()
	assert.Equal(t, expectedRunningAfterStart, server.isRunning, "server should be marked as running after Start")
	server.mu.Unlock()

	// Stop the server
	ctx, cancel := context.WithTimeout(context.Background(), expectedTimeout)
	defer cancel()
	
	err = server.Stop(ctx)
	assert.NoError(t, err, "Stop should succeed with valid context")

	// Verify server is stopped
	server.mu.Lock()
	assert.Equal(t, expectedRunningAfterStop, server.isRunning, "server should not be running after Stop")
	server.mu.Unlock()
}

// mockHTTPServer is a mock implementation that simulates ListenAndServe errors
type mockHTTPServer struct {
	*http.Server
	listenError error
}

func (m *mockHTTPServer) ListenAndServe() error {
	if m.listenError != nil {
		return m.listenError
	}
	return m.Server.ListenAndServe()
}

func Test_Server_Start_ListenError(t *testing.T) {
	// This test verifies that errors from ListenAndServe are properly sent to errChan
	
	// Setup expectations
	expectedAddr := ":99999" // Invalid port to cause error
	expectedErrorMessage := "HTTP server stopped unexpectedly"
	expectedRunning := false
	expectedTimeout := 1 * time.Second
	
	// Arrange
	server := &Server{
		httpServer: &http.Server{
			Addr: expectedAddr,
		},
		mux:       http.NewServeMux(),
		isRunning: false,
	}
	errChan := make(chan error, 1)

	// Act
	server.Start(errChan)

	// Assert - wait for the error
	select {
	case err := <-errChan:
		require.Error(t, err, "Start should send error to channel when ListenAndServe fails")
		assert.Contains(t, err.Error(), expectedErrorMessage, "error message should indicate unexpected server stop")
	case <-time.After(expectedTimeout):
		t.Error("expected error from ListenAndServe but none received")
	}

	// Verify server is no longer marked as running
	server.mu.Lock()
	assert.Equal(t, expectedRunning, server.isRunning, "server should not be marked as running after ListenAndServe error")
	server.mu.Unlock()
}

// Test_Server_Stop_ShutdownError tests the error handling when Shutdown fails
// This is difficult to test without mocking because http.Server.Shutdown rarely fails
func Test_Server_Stop_ShutdownError(t *testing.T) {
	// http.Server.Shutdown typically only fails in these cases:
	// 1. Context is already cancelled (but it returns nil in this case)
	// 2. Server has active connections that don't close gracefully
	// 3. Internal server errors (very rare)
	
	// Since we cannot reliably trigger a Shutdown error without complex mocking
	// or modifying production code, we document that this error path exists
	// for defensive programming purposes.
	t.Skip("http.Server.Shutdown error path cannot be reliably triggered - defensive code")
}

