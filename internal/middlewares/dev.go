package middlewares

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/rs/zerolog"
)

// Simple Middleware used in development to log the request, response and execution duration.
func WithDevLogInterceptor(logger *zerolog.Logger) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()
			logger.Debug().Msgf("BeginAny %s", req.Any())
			logger.Debug().Msgf("Begin %s", req)
			logger.Debug().Msgf("BeginHttp %s", req.HTTPMethod())
			logger.Debug().Msgf("BeginPeer %s", req.Peer())
			logger.Debug().Msgf("BeginSpec %v", req.Spec())
			logger.Debug().Msgf("BeginHeader %s", req.Header())

			res, err := next(ctx, req)
			logger.Debug().Dur("duration", time.Since(start)).Msgf("End %s", req.Any())
			return res, err
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
