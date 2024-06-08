package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kefniark/mango/app/docs/config"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const defaultTimeout = 5 * time.Second

func main() {
	addr := config.AppAddr()

	logger := newLogger()
	logger.Debug().Msg("Initialize Server")

	// Router
	r := chi.NewRouter()

	// Build Default Context (DB, Logger)
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			ctx := config.WithLogger(r.Context(), logger)
			h.ServeHTTP(rw, r.WithContext(ctx))
		})
	})
	registerStaticFilesRoutes(r)
	registerPageRoutes(r)

	// HTTP listen and serve
	logger.Info().Msgf("Listening on %s", config.AppPublicAddr())
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: defaultTimeout,
		Handler:           h2c.NewHandler(r, &http2.Server{}),
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Panic().Err(err).Msgf("Cannot start server and listen on %s", addr)
	}
}

func newLogger() *zerolog.Logger {
	// Pretty Stdout for development
	if config.IsDev() {
		logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
		return &logger
	}

	// Structured JSON Logger for production
	logger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel).With().Timestamp().Caller().Stack().Logger()
	return &logger
}
