package internal

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/kefniark/go-web-server/gen/api/apiconnect"
	"github.com/kefniark/go-web-server/internal/api"
	"github.com/kefniark/go-web-server/internal/core"
	"github.com/kefniark/go-web-server/internal/templates"
)

func registerAPIRoutes(r *chi.Mux, options *core.ServerOptions) {
	r.Route("/api", func(r chi.Router) {
		path, handler := apiconnect.NewUsersHandler(api.NewUserService(options))
		r.Mount(path, http.StripPrefix("/api", handler))

		path, handler = apiconnect.NewProductsHandler(api.NewProductService(options))
		r.Mount(path, http.StripPrefix("/api", handler))
	})
}

func registerStaticFilesRoutes(r *chi.Mux) {
	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
}

func registerPageRoutes(r *chi.Mux, _ *core.ServerOptions) {
	r.Get("/", templ.Handler(templates.Home("Home")).ServeHTTP)
}
