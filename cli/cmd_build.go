package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/kefniark/go-web-server/cli/config"
	"github.com/spf13/cobra"
)

func buildCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Build Mango to a golang binaries",
		Run: func(cmd *cobra.Command, args []string) {
			config.Logger.Info().Msg("Build Apps ...")
			for name, cfg := range *cfg {
				if _, err := os.Stat(path.Join(name, ".mango")); err != nil {
					config.Logger.Warn().Msg("Need preparation first, wait a moment ...")
					for _, exec := range preparer {
						exec.Execute(name)
					}
					for _, exec := range generater {
						exec.Execute(name)
					}
				}

				for _, platform := range cfg.Build.Platforms {
					os.MkdirAll(path.Join("dist", name), os.ModeAppend)

					cmd := exec.Command(
						"env", fmt.Sprintf("GOOS=%s", platform.Os), fmt.Sprintf("GOARCH=%s", platform.Arch),
						"go", "build", "-o", path.Join("dist", name, fmt.Sprintf("%s-%s-%s", name, platform.Os, platform.Arch)), fmt.Sprintf("./%s", name))
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						config.Logger.Error().Err(err).Str("app", name).Msg("Build failed")
					} else {
						config.Logger.Info().Str("app", name).Str("file", path.Join("dist", name, fmt.Sprintf("%s-%s-%s", name, platform.Os, platform.Arch))).Msg("Build Succeed")
					}
				}
			}
			config.Logger.Info().Msg("Build Completed !")
		},
	}
}
