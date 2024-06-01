package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kefniark/mango/docs/config"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const addr = ":5600"
const defaultTimeout = 5 * time.Second

func main() {
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

func newLogger() *zerolog.Logger {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)

	return &logger
}
