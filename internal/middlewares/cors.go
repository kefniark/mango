package middlewares

import (
	"net/http"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

const maxAgeCors = 7200 // 2 hours in seconds

// HTTP Middleware to handle CORS
// ref: https://github.com/connectrpc/cors-go
func CORS(logger *zerolog.Logger) func(handler http.Handler) http.Handler {
	logger.Debug().Msg("Adding CORS middleware")

	return func(connectHandler http.Handler) http.Handler {
		c := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: connectcors.AllowedMethods(),
			AllowedHeaders: connectcors.AllowedHeaders(),
			ExposedHeaders: connectcors.ExposedHeaders(),
			MaxAge:         maxAgeCors,
		})
		return c.Handler(connectHandler)
	}
}
