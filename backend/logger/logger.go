package logger

import (
	"log"
	"log/slog"
	"os"
)

var (
	consoleLogger, fileLogger *slog.Logger
)


// NewLogger initializes the console and file loggers with the specified log level.
func NewLogger(logFile string, level slog.Level) error {
	var err error

	// Open or create the log file with appropriate permissions
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create handlers for console and file output, setting the log level
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	fileHandler := slog.NewTextHandler(file, &slog.HandlerOptions{Level: level})

	// Initialize the loggers with their respective handlers
	consoleLogger = slog.New(consoleHandler)
	fileLogger = slog.New(fileHandler)

	return err
}

func Info(msg string, args ...any) {
	fileLogger.Info(msg, args...)
	consoleLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	fileLogger.Warn(msg, args...)
	consoleLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	fileLogger.Error(msg, args...)
	consoleLogger.Error(msg, args...)
}

func Debug(msg string, args ...any) {
	fileLogger.Debug(msg, args...)
	consoleLogger.Debug(msg, args...)
}
