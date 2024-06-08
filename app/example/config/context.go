package config

import (
	"context"
	"errors"

	"github.com/kefniark/mango/app/example/codegen/database"
	"github.com/rs/zerolog"
)

type contextKey string

const dbContextKey contextKey = "MangoDB"

func GetDB(ctx context.Context) *database.Queries {
	if db, ok := ctx.Value(dbContextKey).(*database.Queries); ok {
		return db
	}
	panic(errors.New("no database in context"))
}

func WithDB(ctx context.Context, db *database.Queries) context.Context {
	return context.WithValue(ctx, dbContextKey, db)
}

const loggerContextKey contextKey = "MangoLogger"

func GetLogger(ctx context.Context) *zerolog.Logger {
	if logger, ok := ctx.Value(loggerContextKey).(*zerolog.Logger); ok {
		return logger
	}
	panic(errors.New("no logger in context"))
}

func WithLogger(ctx context.Context, logger *zerolog.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

const showLayoutContextKey contextKey = "MangoLayout"

func HasLayout(ctx context.Context) bool {
	if val, ok := ctx.Value(showLayoutContextKey).(bool); ok {
		return val
	}
	return true
}

func WithLayout(ctx context.Context, value bool) context.Context {
	return context.WithValue(ctx, showLayoutContextKey, value)
}
