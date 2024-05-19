package internal

import (
	"context"
	"database/sql"
	_ "embed"
	"net/http"
	"os"
	"time"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api/apiconnect"
	dbClient "github.com/kefniark/go-web-server/gen/db"
	productHandlers "github.com/kefniark/go-web-server/internal/api/handlers/products"
	userHandlers "github.com/kefniark/go-web-server/internal/api/handlers/users"
	"github.com/kefniark/go-web-server/internal/middlewares"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	// sqlite3 driver.
	_ "github.com/mattn/go-sqlite3"
)

//go:embed db/schema.sql
var dbSchema string

const addr = "localhost:8081"
const defaultTimeout = 5 * time.Second

func newDB() (*dbClient.Queries, error) {
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

func NewServer() {
	logger := newLogger()
	logger.Debug().Msg("Initialize Server")

	db, err := newDB()
	if err != nil {
		logger.Fatal().Err(err).Msg("Cannot initialize Database connection")
	}

	// middlewares (http) & interceptors (connect)
	middlewareProvider := middlewares.NewMiddlewareProvider()
	middlewareProvider.Add(middlewares.CORS(logger))
	middlewareProvider.Add(middlewares.Auth(db, logger))
	interceptors := connect.WithInterceptors(middlewares.WithDevLogInterceptor(logger))

	// create router
	mux := http.NewServeMux()

	path, handler := apiconnect.NewUsersHandler(userHandlers.NewUserService(db, logger), interceptors)
	mux.Handle(path, handler)

	path, handler = apiconnect.NewProductsHandler(productHandlers.NewProductService(db, logger), interceptors)
	mux.Handle(path, handler)

	// HTTP listen and serve
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: defaultTimeout,
		Handler:           middlewareProvider.Apply(h2c.NewHandler(mux, &http2.Server{})),
	}

	logger.Info().Msgf("Listening on %s", addr)
	if err = server.ListenAndServe(); err != nil {
		logger.Fatal().Err(err).Msgf("Cannot start server and listen on %s", addr)
	}
}
