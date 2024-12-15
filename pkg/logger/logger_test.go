package logger

import (
	"errors"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {

	err := os.Setenv(ScopeLevel, "debug")
	if err != nil {
		t.Fail()
	}
	logger := NewLoggerWithTrace("test-xxxxxxx")

	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message", errors.New("error message"))
	logger.Fatal("This is a fatal message")
	logger.Debug("This is a debug message")

	logger.Infof("User %s logged in successfully", "JohnDoe")
	logger.Debugf("Debugging info: value=%d", 42)
	logger.Warnf("Warning: %s", "You are approaching your limit")
	logger.Errorf("Error: %s", errors.New("error message"), "You are not allowed to do that")
	logger.Fatalf("Fatal error: %s", "The system is shutting down")
}
