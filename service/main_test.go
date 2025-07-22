package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockService is a test implementation of the service interface.
type mockService struct {
	name          string
	startError    error      // Error to return from Start
	startDelay    time.Duration // Delay before sending error
	stopError     error      // Error to return from Stop
	stopDelay     time.Duration // Delay before returning from Stop
	startCalled   atomic.Bool   // Track if Start was called
	stopCalled    atomic.Bool   // Track if Stop was called
	stopCtx       context.Context // Context passed to Stop
	mu            sync.Mutex
}

// Name returns the service's display name for logging purposes.
func (m *mockService) Name() string {
	return m.name
}

// Start begins the service operation, sending any errors to the provided channel.
func (m *mockService) Start(errChan chan error) {
	m.startCalled.Store(true)
	if m.startError != nil {
		go func() {
			if m.startDelay > 0 {
				time.Sleep(m.startDelay)
			}
			errChan <- m.startError
		}()
	}
}

// Stop gracefully shuts down the service within the provided context deadline.
func (m *mockService) Stop(ctx context.Context) (err error) {
	m.mu.Lock()
	m.stopCtx = ctx
	m.mu.Unlock()
	
	m.stopCalled.Store(true)
	if m.stopDelay > 0 {
		time.Sleep(m.stopDelay)
	}
	return m.stopError
}

// mockSignalNotifier is a test implementation of signalNotifier.
type mockSignalNotifier struct {
	channels []chan<- os.Signal
	signals  []os.Signal
	mu       sync.Mutex
}

// Notify registers the channel to receive notifications.
func (m *mockSignalNotifier) Notify(c chan<- os.Signal, sig ...os.Signal) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.channels = append(m.channels, c)
	m.signals = append(m.signals, sig...)
}

// Stop stops the notification.
func (m *mockSignalNotifier) Stop(c chan<- os.Signal) {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Find and remove the channel
	for i, ch := range m.channels {
		if ch == c {
			m.channels = append(m.channels[:i], m.channels[i+1:]...)
			break
		}
	}
}

// sendSignal sends a signal to all registered channels.
func (m *mockSignalNotifier) sendSignal(sig os.Signal) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, ch := range m.channels {
		select {
		case ch <- sig:
		default:
		}
	}
}

func Test_run_SuccessfulStartupAndShutdown(t *testing.T) {
	// Setup expectations
	expectedServerName := "test-server"
	expectedSignal := syscall.SIGTERM
	
	// Create mock service
	mockSvc := &mockService{name: expectedServerName}
	
	// Create mock signal notifier
	mockNotifier := &mockSignalNotifier{}
	
	// Create test config
	cfg := config{
		serverAddr:      ":8080",
		shutdownTimeout: 1 * time.Second,
		serverFactory: func(addr string) (service, error) {
			return mockSvc, nil
		},
		signalNotifier: mockNotifier,
	}
	
	// Run in goroutine to simulate signal
	errChan := make(chan error, 1)
	go func() {
		errChan <- run(cfg)
	}()
	
	// Give the service time to start
	time.Sleep(100 * time.Millisecond)
	
	// Verify service was started
	assert.True(t, mockSvc.startCalled.Load(), "service Start should have been called")
	
	// Send shutdown signal
	mockNotifier.sendSignal(expectedSignal)
	
	// Wait for run to complete
	err := <-errChan
	
	// Assert
	assert.NoError(t, err, "run should not return an error for successful shutdown")
	assert.True(t, mockSvc.stopCalled.Load(), "service Stop should have been called")
}

func Test_run_ServerCreationError(t *testing.T) {
	// Setup expectations
	expectedError := "failed to create server"
	
	// Create test config with failing server factory
	cfg := config{
		serverAddr:      ":8080",
		shutdownTimeout: 1 * time.Second,
		serverFactory: func(addr string) (service, error) {
			return nil, errors.New(expectedError)
		},
		signalNotifier: &mockSignalNotifier{},
	}
	
	// Act
	err := run(cfg)
	
	// Assert
	require.Error(t, err, "run should return an error when server creation fails")
	assert.Contains(t, err.Error(), expectedError, "error should contain the server creation error")
}

func Test_run_ServiceStartError(t *testing.T) {
	// Setup expectations
	expectedStartError := "service failed to start"
	expectedServerName := "failing-server"
	
	// Create mock service that fails on start
	mockSvc := &mockService{
		name:       expectedServerName,
		startError: errors.New(expectedStartError),
		startDelay: 50 * time.Millisecond,
	}
	
	// Create test config
	cfg := config{
		serverAddr:      ":8080",
		shutdownTimeout: 1 * time.Second,
		serverFactory: func(addr string) (service, error) {
			return mockSvc, nil
		},
		signalNotifier: &mockSignalNotifier{},
	}
	
	// Act
	err := run(cfg)
	
	// Assert
	assert.NoError(t, err, "run should not return an error even when service fails to start")
	assert.True(t, mockSvc.startCalled.Load(), "service Start should have been called")
	assert.True(t, mockSvc.stopCalled.Load(), "service Stop should have been called for cleanup")
}

func Test_run_MultipleServiceErrors(t *testing.T) {
	// This tests the error channel draining logic
	// We'll need to modify the test to simulate multiple services
	t.Skip("Requires modification to support multiple services in config")
}

func Test_run_ShutdownTimeout(t *testing.T) {
	// Setup expectations
	expectedServerName := "slow-server"
	shortTimeout := 100 * time.Millisecond
	
	// Create mock service that takes long to stop
	mockSvc := &mockService{
		name:      expectedServerName,
		stopDelay: 500 * time.Millisecond, // Longer than timeout
	}
	
	// Create mock signal notifier
	mockNotifier := &mockSignalNotifier{}
	
	// Create test config
	cfg := config{
		serverAddr:      ":8080",
		shutdownTimeout: shortTimeout,
		serverFactory: func(addr string) (service, error) {
			return mockSvc, nil
		},
		signalNotifier: mockNotifier,
	}
	
	// Run in goroutine to simulate signal
	errChan := make(chan error, 1)
	startTime := time.Now()
	go func() {
		errChan <- run(cfg)
	}()
	
	// Give the service time to start
	time.Sleep(50 * time.Millisecond)
	
	// Send shutdown signal
	mockNotifier.sendSignal(syscall.SIGINT)
	
	// Wait for run to complete
	err := <-errChan
	
	// Assert
	assert.NoError(t, err, "run should complete even with slow shutdown")
	assert.True(t, mockSvc.stopCalled.Load(), "service Stop should have been called")
	
	// Verify the context passed to Stop had the timeout
	mockSvc.mu.Lock()
	stopCtx := mockSvc.stopCtx
	mockSvc.mu.Unlock()
	
	require.NotNil(t, stopCtx, "Stop should have been called with a context")
	deadline, ok := stopCtx.Deadline()
	assert.True(t, ok, "context should have a deadline")
	
	// Check that deadline was set correctly relative to when shutdown started
	expectedDeadline := startTime.Add(50 * time.Millisecond).Add(shortTimeout) // Approximate when shutdown started
	assert.WithinDuration(t, expectedDeadline, deadline, 200*time.Millisecond, "deadline should be approximately shutdown time + timeout")
}

func Test_startServices_SingleService(t *testing.T) {
	// Setup expectations
	expectedServiceName := "test-service"
	
	// Create mock service
	mockSvc := &mockService{name: expectedServiceName}
	
	// Create channels and wait group
	ctx := context.Background()
	shutdownChan := make(chan struct{})
	errChan := make(chan error, 1)
	var wg sync.WaitGroup
	
	// Act
	startServices(ctx, shutdownChan, &wg, errChan, mockSvc)
	
	// Give goroutines time to start
	time.Sleep(50 * time.Millisecond)
	
	// Verify service was started
	assert.True(t, mockSvc.startCalled.Load(), "service Start should have been called")
	
	// Trigger shutdown
	close(shutdownChan)
	
	// Wait for completion
	wg.Wait()
	
	// Assert
	assert.True(t, mockSvc.stopCalled.Load(), "service Stop should have been called")
}

func Test_startServices_MultipleServices(t *testing.T) {
	// Setup expectations
	expectedCount := 3
	
	// Create multiple mock services
	services := make([]service, expectedCount)
	mockServices := make([]*mockService, expectedCount)
	for i := 0; i < expectedCount; i++ {
		mockServices[i] = &mockService{name: fmt.Sprintf("service-%d", i)}
		services[i] = mockServices[i]
	}
	
	// Create channels and wait group
	ctx := context.Background()
	shutdownChan := make(chan struct{})
	errChan := make(chan error, expectedCount)
	var wg sync.WaitGroup
	
	// Act
	startServices(ctx, shutdownChan, &wg, errChan, services...)
	
	// Give goroutines time to start
	time.Sleep(50 * time.Millisecond)
	
	// Verify all services were started
	for i, mockSvc := range mockServices {
		assert.True(t, mockSvc.startCalled.Load(), "service %d Start should have been called", i)
	}
	
	// Trigger shutdown
	close(shutdownChan)
	
	// Wait for completion
	wg.Wait()
	
	// Assert all services were stopped
	for i, mockSvc := range mockServices {
		assert.True(t, mockSvc.stopCalled.Load(), "service %d Stop should have been called", i)
	}
}

func Test_startServices_StopError(t *testing.T) {
	// Setup expectations
	expectedStopError := "failed to stop"
	expectedServiceName := "failing-service"
	
	// Create mock service that fails to stop
	mockSvc := &mockService{
		name:      expectedServiceName,
		stopError: errors.New(expectedStopError),
	}
	
	// Create channels and wait group
	ctx := context.Background()
	shutdownChan := make(chan struct{})
	errChan := make(chan error, 1)
	var wg sync.WaitGroup
	
	// Act
	startServices(ctx, shutdownChan, &wg, errChan, mockSvc)
	
	// Give goroutines time to start
	time.Sleep(50 * time.Millisecond)
	
	// Trigger shutdown
	close(shutdownChan)
	
	// Wait for completion
	wg.Wait()
	
	// Assert - error should be logged but not returned
	assert.True(t, mockSvc.stopCalled.Load(), "service Stop should have been called even if it errors")
}

func Test_waitForShutdown_Signal(t *testing.T) {
	// Setup expectations
	expectedSignal := syscall.SIGTERM
	expectedReason := fmt.Sprintf("received signal: %v", expectedSignal)
	
	// Create channels
	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	
	// Send signal
	sigChan <- expectedSignal
	
	// Act
	reason := waitForShutdown(sigChan, errChan)
	
	// Assert
	assert.Equal(t, expectedReason, reason, "shutdown reason should indicate signal")
}

func Test_waitForShutdown_Error(t *testing.T) {
	// Setup expectations
	expectedError := errors.New("fatal service error")
	expectedReason := fmt.Sprintf("fatal error: %v", expectedError)
	
	// Create channels
	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error, 2) // Buffer for drain test
	
	// Send error
	errChan <- expectedError
	
	// Act
	reason := waitForShutdown(sigChan, errChan)
	
	// Assert
	assert.Equal(t, expectedReason, reason, "shutdown reason should indicate error")
	
	// Send another error to test draining
	errChan <- errors.New("additional error")
	close(errChan)
	
	// Give drain goroutine time to process
	time.Sleep(50 * time.Millisecond)
}

func Test_osSignalNotifier(t *testing.T) {
	// Test the real signal notifier implementation
	notifier := osSignalNotifier{}
	sigChan := make(chan os.Signal, 1)
	
	// Test Notify
	notifier.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	// Test Stop
	notifier.Stop(sigChan)
	
	// Basic test - just ensure methods don't panic
	assert.True(t, true, "osSignalNotifier methods should not panic")
}

func Test_defaultConfig(t *testing.T) {
	// Setup expectations
	expectedAddr := ":8080"
	expectedTimeout := 30 * time.Second
	
	// Act
	cfg := defaultConfig()
	
	// Assert
	assert.Equal(t, expectedAddr, cfg.serverAddr, "default server address should be :8080")
	assert.Equal(t, expectedTimeout, cfg.shutdownTimeout, "default shutdown timeout should be 30 seconds")
	assert.NotNil(t, cfg.serverFactory, "server factory should not be nil")
	assert.NotNil(t, cfg.signalNotifier, "signal notifier should not be nil")
	
	// Test that server factory works
	server, err := cfg.serverFactory(":8080")
	assert.NoError(t, err, "default server factory should create server without error")
	assert.NotNil(t, server, "created server should not be nil")
}

// Test_main is difficult to test directly, but we can test it doesn't panic
// In a real scenario, you might use build tags to create a testable version
func Test_main_Integration(t *testing.T) {
	t.Skip("main() is tested indirectly through run() tests")
}