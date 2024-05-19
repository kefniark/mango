package middlewares

import (
	"context"
	"net/http"
	"strings"

	authn "connectrpc.com/authn"
	dbClient "github.com/kefniark/go-web-server/gen/db"
	"github.com/rs/zerolog"
)

func Auth(db *dbClient.Queries, logger *zerolog.Logger) func(handler http.Handler) http.Handler {
	logger.Debug().Msg("Adding Auth middleware")
	auth := authenticate(db, logger)

	return func(handler http.Handler) http.Handler {
		return authn.NewMiddleware(auth).Wrap(handler)
	}
}

func authenticate(_ *dbClient.Queries, logger *zerolog.Logger) func(_ context.Context, req authn.Request) (any, error) {
	return func(_ context.Context, req authn.Request) (any, error) {
		token, ok := bearerToken(req)
		if !ok || token != "none" {
			logger.Debug().Str("token", token).Msg("Authenticating failed")
			// return nil, authn.Errorf("invalid password")
		}

		logger.Debug().Str("token", token).Msg("Authenticating succeeded")
		return token, nil
	}
}

func bearerToken(req authn.Request) (string, bool) {
	auth := req.Header().Get("Authorization")
	const prefix = "Bearer "
	if !strings.HasPrefix(auth, prefix) {
		return "", false
	}
	return auth[len(prefix):], true
}
