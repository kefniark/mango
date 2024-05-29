package prepare

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/kefniark/mango/cli/config"
)

type NodeJSPrepare struct{}

func (prepare NodeJSPrepare) Name() string {
	return "NodeJS Preparer"
}

func (prepare NodeJSPrepare) Execute(app string) error {
	if _, err := os.Stat(path.Join(app, "tools", "nodejs", "package.json")); err != nil {
		config.Logger.Info().Str("app", app).Str("path", path.Join(app, "tools", "nodejs")).Msg("Skip Install Nodejs ...")
		return nil
	}

	dir, err := filepath.Abs(path.Join(app, "tools", "nodejs"))
	if err != nil {
		return err
	}

	cmd := exec.Command("npm", "install")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	config.Logger.Info().Str("app", app).Str("path", path.Join(app, "tools", "nodejs")).Msg("Installed Nodejs")
	return nil
}
