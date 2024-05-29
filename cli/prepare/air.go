package prepare

import (
	"os"
	"path"
	"text/template"

	"github.com/kefniark/mango/cli/config"
)

type AirConfig struct {
	App       string
	AppPort   int
	ProxyPort int
}

type AirPrepare struct{}

func (prepare AirPrepare) Name() string {
	return "AIR Preparer"
}

func (prepare AirPrepare) Execute(app string) error {
	tmpl, err := template.ParseFS(templates, "templates/air.go.tmpl")
	if err != nil {
		return err
	}

	folderMango := path.Join(app, ".mango")
	err = os.MkdirAll(folderMango, os.ModePerm)
	if err != nil {
		return err
	}

	config.Logger.Debug().Str("app", app).Str("path", path.Join(folderMango, "air.toml")).Msg("Generate Air File")
	f, err := os.Create(path.Join(folderMango, "air.toml"))
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, AirConfig{
		App:       app,
		AppPort:   5600,
		ProxyPort: 5500,
	})

	return err
}
