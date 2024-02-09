package logger

import (
	"context"
	"log/slog"
	"os"
	"sync"
)

var (
	theLog   *slog.Logger
	logMutex sync.Mutex
)

type Logger interface {
	DebugContext(ctx context.Context, msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
}

// InitLogger to initialize logger
func InitLogger(level slog.Level) {
	logMutex.Lock()
	defer logMutex.Unlock()

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}

	handler := ContextHandler{slog.NewJSONHandler(os.Stdout, opts)}
	theLog = slog.New(handler)
}

// GetLogger to get context logger
func GetLogger() Logger {
	if theLog == nil {
		panic("logger is nil")
	}

	return theLog
}
