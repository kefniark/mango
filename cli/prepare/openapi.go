package prepare

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/kefniark/go-web-server/cli/config"
)

type OpenAPIPrepare struct {
	Config *config.Config
}

func (prepare OpenAPIPrepare) Execute(app string) error {
	tmpl, err := template.ParseFS(templates, "templates/openapi-config.go.tmpl")
	if err != nil {
		return err
	}

	folderAPI := path.Join(app, ".mango/api")
	err = os.MkdirAll(folderAPI, os.ModeAppend)
	if err != nil {
		return err
	}

	config.Logger.Debug().Str("app", app).Str("path", path.Join(folderAPI, "openapi.yaml")).Msg("Generated File")
	f, err := os.Create(path.Join(folderAPI, "openapi.yaml"))
	if err != nil {
		return err
	}

	config, ok := (*prepare.Config)[app]
	if !ok {
		return fmt.Errorf("cannot find config %s", app)
	}

	return tmpl.Execute(f, config)
}
