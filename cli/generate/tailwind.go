package generate

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/kefniark/mango/cli/config"
)

type TailwindGenerator struct{}

func (generator TailwindGenerator) Execute(app string) error {
	if _, err := os.Stat(path.Join(app, "views")); err != nil {
		config.Logger.Debug().Str("app", app).Str("path", path.Join(app, "views")).Msg("Skip Generate Tailwind ...")
		return nil
	}

	dir, err := filepath.Abs(path.Join(app, "tools", "nodejs"))
	if err != nil {
		return err
	}

	cmd := exec.Command("npx", "tailwindcss", "-i", "./tailwind.css", "-o", "../../assets/css/tailwind.css")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	config.Logger.Debug().Str("app", app).Str("path", path.Join(app, "views")).Msg("Generate Templ")

	return nil
}
