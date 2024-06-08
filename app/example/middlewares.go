package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/kefniark/mango/app/example/config"
)

const corsMaxAge = 3600

func registerMiddlewares(r *chi.Mux) {
	l := httplog.NewLogger("app", httplog.Options{JSON: !config.IsDev()})

	r.Use(skipLayoutMiddleware)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(httplog.RequestLogger(l))
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(middleware.Timeout(defaultTimeout))

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Layout"},
		ExposedHeaders:   []string{""},
		AllowCredentials: false,
		MaxAge:           corsMaxAge,
	}))
}

func skipLayoutMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		val := r.Header.Get("Layout")
		ctx := config.WithLayout(r.Context(), val != "false")
		h.ServeHTTP(rw, r.WithContext(ctx))
	})
}
