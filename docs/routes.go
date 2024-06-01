package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/kefniark/mango/docs/config"
	_ "github.com/kefniark/mango/docs/views/pages"
)

const assetCache = 3600 * 24

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

func registerPageRoutes(r *chi.Mux) {
	config.Register(r)
}
