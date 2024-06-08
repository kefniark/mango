package config

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
)

type contextKey string

const loggerContextKey contextKey = "MangoLogger"

func GetLogger(ctx context.Context) *zerolog.Logger {
	if logger, ok := ctx.Value(loggerContextKey).(*zerolog.Logger); ok {
		return logger
	}
	panic(errors.New("no logger in context"))
}

func WithLogger(ctx context.Context, db *zerolog.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, db)
}
