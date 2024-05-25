package internal

import (
	"net/http"
	"time"

	"github.com/kefniark/go-web-server/internal/middlewares"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const addr = "localhost:8081"
const defaultTimeout = 5 * time.Second

func NewServer() {
	options := newServerOptions()
	logger := options.Logger
	logger.Debug().Msg("Initialize Server")

	// middlewares (http) & interceptors (connect)
	middlewareProvider := middlewares.NewMiddlewareProvider()
	middlewareProvider.Add(middlewares.CORS(options))
	middlewareProvider.Add(middlewares.Auth(options))

	// create router
	mux := http.NewServeMux()

	registerRoutes(mux, options)

	// HTTP listen and serve
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: defaultTimeout,
		Handler:           middlewareProvider.Apply(h2c.NewHandler(mux, &http2.Server{})),
	}

	logger.Info().Msgf("Listening on %s", addr)
	if err := server.ListenAndServe(); err != nil {
		logger.Panic().Err(err).Msgf("Cannot start server and listen on %s", addr)
	}
}
