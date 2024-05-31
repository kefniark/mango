package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/kefniark/mango/example/codegen/database"
	"github.com/kefniark/mango/example/config"
	"github.com/kefniark/mango/example/db"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const addr = ":5600"
const defaultTimeout = 5 * time.Second

func main() {
	logger, db := newServerOptions()
	logger.Debug().Msg("Initialize Server")

	// Router
	r := chi.NewRouter()

	// Build Default Context (DB, Logger)
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			ctx := config.WithDB(r.Context(), db)
			ctx = config.WithLogger(ctx, logger)
			h.ServeHTTP(rw, r.WithContext(ctx))
		})
	})
	registerMiddlewares(r)
	registerAPIRoutes(r)
	registerStaticFilesRoutes(r)
	registerPageRoutes(r)

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

func newServerOptions() (*zerolog.Logger, *database.Queries) {
	logger := newLogger()

	dbClient, err := db.New()
	if err != nil {
		logger.Panic().Err(err).Msg("Cannot initialize Database connection")
	}

	if err = dbClient.Migrate(); err != nil {
		logger.Warn().Err(err).Msg("Cannot apply database schema")
	} else {
		client := dbClient.Client()
		for i := range 5 {
			_, err = client.SetUser(context.Background(), database.SetUserParams{
				ID:   uuid.New(),
				Name: fmt.Sprintf("user-%d", i),
				Bio:  "",
			})
			if err != nil {
				logger.Warn().Err(err).Msg("Cannot seed data")
			}
		}
	}

	return logger, dbClient.Client()
}

func newLogger() *zerolog.Logger {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)

	return &logger
}
