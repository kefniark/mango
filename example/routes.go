package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/kefniark/go-web-server/example/codegen/api/apiconnect"

	"github.com/kefniark/go-web-server/example/api/handlers"
	"github.com/kefniark/go-web-server/example/config"
	"github.com/kefniark/go-web-server/example/views/pages"
)

func registerAPIRoutes(r *chi.Mux, options *config.ServerOptions) {
	r.Route("/api", func(r chi.Router) {
		path, handler := apiconnect.NewUsersHandler(handlers.NewUserService(options))
		r.Mount(path, http.StripPrefix("/api", handler))

		path, handler = apiconnect.NewProductsHandler(handlers.NewProductService(options))
		r.Mount(path, http.StripPrefix("/api", handler))
	})
}

func registerStaticFilesRoutes(r *chi.Mux) {
	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
}

func registerPageRoutes(r *chi.Mux, _ *config.ServerOptions) {
	r.Get("/", templ.Handler(pages.Home("Home")).ServeHTTP)
	r.Get("/about", templ.Handler(pages.About("About")).ServeHTTP)
}
