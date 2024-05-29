package generate

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/kefniark/mango/cli/config"
)

type SQLCGenerator struct{}

func (prepare SQLCGenerator) Name() string {
	return "SQLC Generator"
}

func (generator SQLCGenerator) Execute(app string) error {
	if _, err := os.Stat(path.Join(app, "db")); err != nil {
		config.Logger.Debug().Str("app", app).Str("path", path.Join(app, "codegen", "db")).Msg("Skip Generate DB ...")
		return nil
	}

	dir, err := filepath.Abs(app)
	if err != nil {
		return err
	}

	sqlcConfig := path.Join(dir, ".mango", "db", "sqlc.yaml")

	config.Logger.Debug().Str("path", sqlcConfig).Msg("Is file exist")

	cmd := exec.Command("sqlc", "generate", "-f", sqlcConfig)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	config.Logger.Debug().Str("app", app).Str("path", path.Join(app, "codegen", "db")).Msg("Generated DB")
	return nil
}
