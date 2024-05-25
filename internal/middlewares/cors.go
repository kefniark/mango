package middlewares

import (
	"net/http"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"

	"github.com/kefniark/go-web-server/internal/core"
)

const maxAgeCors = 7200 // 2 hours in seconds

// HTTP Middleware to handle CORS
// ref: https://github.com/connectrpc/cors-go
func CORS(options *core.ServerOptions) func(handler http.Handler) http.Handler {
	options.Logger.Debug().Msg("Adding CORS middleware")

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
