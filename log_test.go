package i18n

import (
	"context"
	"testing"
)

// mockLogger is a simple implementation of the Logger interface for testing
type mockLogger struct {
	debugCalled   bool
	errorCalled   bool
	warningCalled bool
	lastFormat    string
	lastArgs      []any
}

func (m *mockLogger) Debugf(_ context.Context, format string, args ...any) {
	m.debugCalled = true
	m.lastFormat = format
	m.lastArgs = args
}

func (m *mockLogger) Errorf(_ context.Context, format string, args ...any) {
	m.errorCalled = true
	m.lastFormat = format
	m.lastArgs = args
}

func (m *mockLogger) Warningf(_ context.Context, format string, args ...any) {
	m.warningCalled = true
	m.lastFormat = format
	m.lastArgs = args
}

func TestLoggingFunctions(t *testing.T) {
	// Save the original logger
	originalLogger := log
	defer func() {
		// Restore the original logger after the test
		log = originalLogger
	}()

	// Create a mock logger
	mock := &mockLogger{}
	log = mock

	ctx := context.Background()
	format := "test %s"
	args := []any{"message"}

	// Test debugf
	debugf(ctx, format, args...)
	if !mock.debugCalled {
		t.Error("Expected debugf to call log.Debugf")
	}
	if mock.lastFormat != format {
		t.Errorf("Expected format to be %q, got %q", format, mock.lastFormat)
	}
	if len(mock.lastArgs) != len(args) || mock.lastArgs[0] != args[0] {
		t.Errorf("Expected args to be %v, got %v", args, mock.lastArgs)
	}

	// Reset mock
	mock = &mockLogger{}
	log = mock

	// Test errorf
	errorf(ctx, format, args...)
	if !mock.errorCalled {
		t.Error("Expected errorf to call log.Errorf")
	}
	if mock.lastFormat != format {
		t.Errorf("Expected format to be %q, got %q", format, mock.lastFormat)
	}
	if len(mock.lastArgs) != len(args) || mock.lastArgs[0] != args[0] {
		t.Errorf("Expected args to be %v, got %v", args, mock.lastArgs)
	}

	// Reset mock
	mock = &mockLogger{}
	log = mock

	// Test warningf
	warningf(ctx, format, args...)
	if !mock.warningCalled {
		t.Error("Expected warningf to call log.Warningf")
	}
	if mock.lastFormat != format {
		t.Errorf("Expected format to be %q, got %q", format, mock.lastFormat)
	}
	if len(mock.lastArgs) != len(args) || mock.lastArgs[0] != args[0] {
		t.Errorf("Expected args to be %v, got %v", args, mock.lastArgs)
	}

	// Test with nil logger
	log = nil
	// These should not panic
	debugf(ctx, format, args...)
	errorf(ctx, format, args...)
	warningf(ctx, format, args...)
}
