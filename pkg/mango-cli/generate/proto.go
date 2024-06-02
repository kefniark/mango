package generate

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/kefniark/mango/pkg/mango-cli/config"
)

type ProtoGenerator struct{}

func (prepare ProtoGenerator) Name() string {
	return "Proto Generator"
}

func (generator ProtoGenerator) Execute(app string) error {
	if _, err := os.Stat(path.Join(app, "api")); err != nil {
		config.Logger.Info().Str("app", app).Str("path", path.Join(app, "codegen", "api")).Msg("Skip Generate API ...")
		return nil
	}

	out, err := exec.Command("which", "buf").Output()
	if err != nil {
		config.Logger.Err(err).Msg("cannot find buf")
		return err
	}

	bufCmd := strings.TrimSpace(string(out))
	if len(bufCmd) == 0 {
		return errors.New("cannot find buf")
	}

	if err := preProtoCodegen(app); err != nil {
		return err
	}

	if err := execProtoCodegen(app, bufCmd); err != nil {
		return err
	}

	if err := postProtoCodegen(app); err != nil {
		return err
	}

	return nil
}

func preProtoCodegen(app string) error {
	entries, err := os.ReadDir(path.Join(app, ".mango", "proto"))
	if err != nil {
		return err
	}

	for _, f := range entries {
		src, err := os.Open(path.Join(app, ".mango", "proto", f.Name()))
		if err != nil {
			return err
		}

		dest, err := os.Create(path.Join(app, f.Name()))
		if err != nil {
			return err
		}

		_, err = io.Copy(dest, src)
		if err != nil {
			return err
		}
	}

	return nil
}

func execProtoCodegen(app string, bufCmd string) error {
	dir, err := filepath.Abs(app)
	if err != nil {
		config.Logger.Err(err).Msg("cannot find app")
		return err
	}

	cmd := exec.Command(bufCmd, "generate")

	config.Logger.Debug().Str("app", app).Str("bufCmd", bufCmd).Str("dir", dir).Msg("Run Buf")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		config.Logger.Err(err).Msg("cannot generate api")
		return err
	}

	config.Logger.Debug().Str("app", app).Str("path", path.Join(app, "codegen", "api")).Msg("Generated API")
	return nil
}

func postProtoCodegen(app string) error {
	entries, err := os.ReadDir(path.Join(app, ".mango", "proto"))
	if err != nil {
		config.Logger.Err(err).Msg("cannot check proto")
		return err
	}

	for _, f := range entries {
		if err = os.Remove(path.Join(app, f.Name())); err != nil {
			return err
		}
	}

	return nil
}
