package prepare

import (
	"os"
	"path"
	"text/template"

	"github.com/kefniark/mango/cli/config"
)

type StaticBuildPrepare struct {
	Config *config.Config
}

func (prepare StaticBuildPrepare) Name() string {
	return "Static Preparer"
}

func (prepare StaticBuildPrepare) Execute(app string) error {
	tmpl, err := template.ParseFS(templates, "templates/static-build.go.tmpl")
	if err != nil {
		return err
	}

	folderDB := path.Join(app, ".mango/go")
	if err = os.MkdirAll(folderDB, os.ModePerm); err != nil {
		return err
	}

	config.Logger.Debug().Str("app", app).Str("path", path.Join(folderDB, "builder.go")).Msg("Generated File")
	f, err := os.Create(path.Join(folderDB, "builder.go"))
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, StaticArgs{Name: app})
}

type StaticArgs struct {
	Name string
}
