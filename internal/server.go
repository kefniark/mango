package internal

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	dbClient "github.com/kefniark/go-web-server/gen/db"
	"github.com/kefniark/go-web-server/internal/core"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	// embed database schema.
	_ "embed"

	// sqlite3 driver.
	_ "github.com/mattn/go-sqlite3"
)

const addr = ":5550"
const defaultTimeout = 5 * time.Second

func NewServer() {
	options := newServerOptions()
	logger := options.Logger
	logger.Debug().Msg("Initialize Server")

	// Router
	r := chi.NewRouter()
	registerMiddlewares(r)
	registerAPIRoutes(r, options)
	registerStaticFilesRoutes(r)
	registerPageRoutes(r, options)

	// HTTP listen and serve
	logger.Info().Msgf("Listening on %s", addr)
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: defaultTimeout,
		Handler:           h2c.NewHandler(r, &http2.Server{}),
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Panic().Err(err).Msgf("Cannot start server and listen on %s", addr)
	}
}

//go:embed db/schema.sql
var dbSchema string

func newServerOptions() *core.ServerOptions {
	logger := newLogger()

	db, err := newDatabase(logger)
	if err != nil {
		logger.Panic().Err(err).Msg("Cannot initialize Database connection")
	}

	return &core.ServerOptions{Logger: logger, DB: db}
}

func newDatabase(logger *zerolog.Logger) (*dbClient.Queries, error) {
	// db, err := sql.Open("sqlite3", ":memory:")
	db, err := sql.Open("sqlite3", "dev.db")
	if err != nil {
		return nil, err
	}

	// apply db schema
	if _, err = db.ExecContext(context.Background(), dbSchema); err != nil {
		logger.Warn().Err(err).Msg("Cannot apply database schema")
	}

	return dbClient.New(db), nil
}

func newLogger() *zerolog.Logger {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)

	return &logger
}
