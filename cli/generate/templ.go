package generate

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/kefniark/mango/cli/config"
)

type TemplGenerator struct{}

func (prepare TemplGenerator) Name() string {
	return "Templ Generator"
}

func (generator TemplGenerator) Execute(app string) error {
	if _, err := os.Stat(path.Join(app, "views")); err != nil {
		config.Logger.Debug().Str("app", app).Str("path", path.Join(app, "views")).Msg("Skip Generate Templ ...")
		return nil
	}

	dir, err := filepath.Abs(app)
	if err != nil {
		return err
	}

	cmd := exec.Command("templ", "generate")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	config.Logger.Debug().Str("app", app).Str("path", path.Join(app, "views")).Msg("Generate Templ")

	return nil
}
