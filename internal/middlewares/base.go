package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type MiddlewareProvider struct {
	middlewares []Middleware
}

func NewMiddlewareProvider() *MiddlewareProvider {
	return &MiddlewareProvider{}
}

func (provider *MiddlewareProvider) Add(middleware Middleware) {
	provider.middlewares = append(provider.middlewares, middleware)
}

func (provider MiddlewareProvider) Apply(handler http.Handler) http.Handler {
	current := handler
	for _, m := range provider.middlewares {
		current = m(current)
	}
	return current
}
