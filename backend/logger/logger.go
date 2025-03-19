package logger

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"time"
)

type CustomHandler struct {
	writer *os.File
	level  slog.Level
}

func (h *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	timestamp := time.Now().Format(time.RFC3339)

	logEntry := map[string]interface{}{
		"level": r.Level.String(),
		"msg":   r.Message,
		"time":  timestamp,
	}

	logData, err := json.Marshal(logEntry)
	if err != nil {
		return err
	}

	h.writer.Write(append(logData, '\n'))
	os.Stdout.Write(append(logData, '\n'))
	return nil
}

func (h *CustomHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *CustomHandler) WithGroup(name string) slog.Handler {
	return h
}

// NewLogger initializes structured logging with the custom handler
func NewLogger(logFile string, level slog.Level) error {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return err
	}

	handler := &CustomHandler{writer: file, level: level}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return nil
}

func Info(msg string, args ...any)  { slog.Info(msg, args...) }
func Warn(msg string, args ...any)  { slog.Warn(msg, args...) }
func Error(msg string, args ...any) { slog.Error(msg, args...) }
func Debug(msg string, args ...any) { slog.Debug(msg, args...) }
