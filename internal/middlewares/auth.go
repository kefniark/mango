package middlewares

import (
	"context"
	"net/http"
	"strings"

	authn "connectrpc.com/authn"
	"github.com/kefniark/go-web-server/internal/core"
)

func Auth(options *core.ServerOptions) func(handler http.Handler) http.Handler {
	options.Logger.Debug().Msg("Adding Auth middleware")
	auth := authenticate(options)

	return func(handler http.Handler) http.Handler {
		return authn.NewMiddleware(auth).Wrap(handler)
	}
}

func authenticate(_ *core.ServerOptions) func(_ context.Context, req authn.Request) (any, error) {
	return func(ctx context.Context, req authn.Request) (any, error) {
		token, ok := bearerToken(req)
		if ok && token != "none" {
			return core.NewUserInfo(token, "MyUser"), nil
		}

		return nil, nil
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
