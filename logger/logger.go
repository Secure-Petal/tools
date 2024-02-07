package logger

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

// InitLogger to initialize logger
func InitLogger() {
	logMutex.Lock()
	defer logMutex.Unlock()

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     getLevel(os.Getenv("LOG_LEVEL")),
	}

	handler := ContextHandler{slog.NewJSONHandler(os.Stdout, opts)}
	theLog = &Logger{slog.New(handler)}
}

func getLevel(l string) slog.Level {
	switch l {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelWarn
	}
}

// GetLogger to get context logger
func GetLogger() *Logger {
	if theLog != nil {
		return theLog
	}

	InitLogger()

	return theLog
}
