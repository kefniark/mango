package internal

import (
	"context"
	"database/sql"
	"os"

	dbClient "github.com/kefniark/go-web-server/gen/db"
	"github.com/kefniark/go-web-server/internal/core"
	"github.com/rs/zerolog"

	// embed database schema.
	_ "embed"

	// sqlite3 driver.
	_ "github.com/mattn/go-sqlite3"
)

//go:embed db/schema.sql
var dbSchema string

func newServerOptions() *core.ServerOptions {
	logger := newLogger()

	db, err := newDatabase()
	if err != nil {
		logger.Panic().Err(err).Msg("Cannot initialize Database connection")
	}

	return &core.ServerOptions{
		Logger: logger,
		DB:     db,
	}
}

func newDatabase() (*dbClient.Queries, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	if _, err = db.ExecContext(context.Background(), dbSchema); err != nil {
		return nil, err
	}

	return dbClient.New(db), nil
}

func newLogger() *zerolog.Logger {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)

	return &logger
}
