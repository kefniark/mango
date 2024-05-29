package prepare

import (
	"embed"
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/kefniark/mango/cli/config"
)

//go:embed templates/*.tmpl
var templates embed.FS

type SQLCPrepare struct {
	Config *config.Config
}

func (prepare SQLCPrepare) Name() string {
	return "SQLC Preparer"
}

func (prepare SQLCPrepare) Execute(app string) error {
	tmpl, err := template.ParseFS(templates, "templates/sqlc-config.go.tmpl")
	if err != nil {
		return err
	}

	folderDB := path.Join(app, ".mango/db")
	if err = os.MkdirAll(folderDB, os.ModePerm); err != nil {
		return err
	}

	config.Logger.Debug().Str("app", app).Str("path", path.Join(folderDB, "sqlc.yaml")).Msg("Generated File")
	f, err := os.Create(path.Join(folderDB, "sqlc.yaml"))
	if err != nil {
		return err
	}

	config, ok := (*prepare.Config)[app]
	if !ok {
		return fmt.Errorf("cannot find config %s", app)
	}

	return tmpl.Execute(f, config.Database)
}
