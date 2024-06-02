package prepare

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/kefniark/mango/pkg/mango-cli/config"
)

type DBPrepare struct {
	Config *config.Config
}

func (prepare DBPrepare) Name() string {
	return "AIR Preparer"
}

func (prepare DBPrepare) getDBTemplate(cfg config.ConfigApp) (*template.Template, error) {
	if strings.TrimSpace(cfg.Database.Engine) == "" {
		config.Logger.Debug().Str("app", cfg.Title).Msg("Skip Generating DB file")
		return nil, nil
	}

	if cfg.Database.Engine == "postgresql" {
		return template.ParseFS(templates, "templates/db-postgres.go.tmpl")
	}

	if cfg.Database.Engine == "sqlite" {
		return template.ParseFS(templates, "templates/db-sqlite.go.tmpl")
	}

	return nil, fmt.Errorf("unknown db %s", cfg.Database.Engine)
}

func (prepare DBPrepare) Execute(app string) error {
	cfg, ok := (*prepare.Config)[app]
	if !ok {
		return fmt.Errorf("cannot find config %s", app)
	}

	tmpl, err := prepare.getDBTemplate(cfg)
	if err != nil {
		return err
	}
	if tmpl == nil {
		return nil
	}

	f, err := os.Create(path.Join(app, "db", "main.go"))
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, cfg.Database)

	return err
}
