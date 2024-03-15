package logger

import (
	"context"
	"sync"
)

var (
	logger     Logger
	loggerOnce sync.Once
)

type Logger interface {
	Instance() any
	Debug(ctx context.Context, msg string, fields ...any)
	Info(ctx context.Context, msg string, fields ...any)
	Warn(ctx context.Context, msg string, fields ...any)
	Error(ctx context.Context, msg string, fields ...any)
	Fatal(ctx context.Context, msg string, fields ...any)
}

func Init(l Logger) Logger {
	loggerOnce.Do(func() {
		logger = l
	})
	return logger
}

func Debug(ctx context.Context, msg string, fields ...any) {
	logger.Debug(ctx, msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...any) {
	logger.Info(ctx, msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...any) {
	logger.Warn(ctx, msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...any) {
	logger.Error(ctx, msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...any) {
	logger.Fatal(ctx, msg, fields...)
}
