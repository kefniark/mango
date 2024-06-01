package config

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type Route struct {
	Pattern string
	Handler *templ.ComponentHandler
}

//nolint:gochecknoglobals // Use for code-generation, not runtime
var Routes []Route = []Route{}

func RegisterPage(pattern string, handler *templ.ComponentHandler) {
	Routes = append(Routes, Route{Pattern: pattern, Handler: handler})
}

func Register(r *chi.Mux) {
	for _, route := range Routes {
		r.Get(route.Pattern, route.Handler.ServeHTTP)
	}
}
