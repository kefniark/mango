package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/kefniark/go-web-server/example/codegen/api/apiconnect"

	"github.com/kefniark/go-web-server/example/api/handlers"
	"github.com/kefniark/go-web-server/example/config"
	"github.com/kefniark/go-web-server/example/views/pages"
)

const assetCache = 3600 * 24

func registerAPIRoutes(r *chi.Mux, options *config.ServerOptions) {
	r.Route("/api", func(r chi.Router) {
		path, handler := apiconnect.NewUsersHandler(handlers.NewUserService(options))
		r.Mount(path, http.StripPrefix("/api", handler))

		path, handler = apiconnect.NewProductsHandler(handlers.NewProductService(options))
		r.Mount(path, http.StripPrefix("/api", handler))
	})
}

//go:embed assets
var embeddedAssetsFS embed.FS

func registerStaticFilesRoutes(r *chi.Mux) {
	assetsFs, err := fs.Sub(embeddedAssetsFS, "assets")
	if err != nil {
		log.Fatal(err)
	}

	fileserver := http.StripPrefix("/assets", http.FileServer(http.FS(assetsFs)))

	r.HandleFunc("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", assetCache))
		fileserver.ServeHTTP(w, r)
	})
}

func registerPageRoutes(r *chi.Mux, _ *config.ServerOptions) {
	r.Get("/", templ.Handler(pages.Home("Home")).ServeHTTP)
	r.Get("/about", templ.Handler(pages.About("About")).ServeHTTP)
}
