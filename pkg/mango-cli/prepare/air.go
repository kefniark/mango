package prepare

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/kefniark/mango/pkg/mango-cli/config"
)

type AirConfig struct {
	App       string
	AppPort   int
	ProxyPort int
}

type AirPrepare struct {
	Config *config.Config
}

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
	defer f.Close()

	cfg, ok := (*prepare.Config)[app]
	if !ok {
		return fmt.Errorf("cannot find config %s", app)
	}

	err = tmpl.Execute(f, AirConfig{
		App:       app,
		AppPort:   cfg.Port + 10000,
		ProxyPort: cfg.Port,
	})

	return err
}
