package logger

import (
	"context"

	"go.uber.org/zap"
)

type correlationIdType int

const (
	CtxRequestIDKey = "X-Request-Id"
)

var logger *zap.Logger

func init() {
	// a fallback/root logger for events without context
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	logger, _ = config.Build()
}

// WithRqId returns a context which knows its request ID
func WithRqId(ctx context.Context, rid string) context.Context {
	return context.WithValue(ctx, CtxRequestIDKey, rid)
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) *zap.Logger {
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(CtxRequestIDKey).(string); ok {
			return logger.With(zap.String("rid", ctxRqId))
		}
	}
	return logger
}
