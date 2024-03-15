package logger

import (
	"context"

	"github.com/hibiken/asynq"
)

func ConvertLoggerToAsynq(oldLogger Logger) asynq.Logger {
	return &loggerConverter{oldLogger}
}

type loggerConverter struct {
	oldLogger Logger
}

func (lc *loggerConverter) Debug(args ...interface{}) {
	// Discard the context and fields
	msg := args[0].(string)
	lc.oldLogger.Debug(context.Background(), msg)
}

func (lc *loggerConverter) Info(args ...interface{}) {
	// Discard the context and fields
	msg := args[0].(string)
	lc.oldLogger.Info(context.Background(), msg)
}

func (lc *loggerConverter) Warn(args ...interface{}) {
	// Discard the context and fields
	msg := args[0].(string)
	lc.oldLogger.Warn(context.Background(), msg)
}

func (lc *loggerConverter) Error(args ...interface{}) {
	// Discard the context and fields
	msg := args[0].(string)
	lc.oldLogger.Error(context.Background(), msg)
}

func (lc *loggerConverter) Fatal(args ...interface{}) {
	// Discard the context and fields
	msg := args[0].(string)
	lc.oldLogger.Fatal(context.Background(), msg)
}
