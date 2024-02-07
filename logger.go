package tools

import (
	"log/slog"
	"os"
	"sync"
)

var (
	theLog   *Logger
	logMutex sync.Mutex
)

// Logger for our own custom logging over the slog package
type Logger struct {
	log *slog.Logger
}

// isInitialized to check if logger is not nil
func isInitialized() bool {
	logMutex.Lock()
	defer logMutex.Unlock()

	return theLog != nil
}

// initLogger to initialize logger
func initLogger() {
	logMutex.Lock()
	defer logMutex.Unlock()

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	theLog = &Logger{slog.New(handler)}
}

// GetLogger to get context logger
func GetLogger() *Logger {
	if isInitialized() {
		return theLog
	}

	initLogger()

	return theLog
}
