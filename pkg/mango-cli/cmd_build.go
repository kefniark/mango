package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/kefniark/mango/pkg/mango-cli/config"
	"github.com/spf13/cobra"
)

func buildCmd(cfg *config.Config) *cobra.Command {
	var filter string
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build Mango to a golang binaries",
		Run: func(cmd *cobra.Command, args []string) {
			config.Logger.Info().Msg("Build Apps ...")
			for name, cfg := range *cfg {
				if filter != "" && filter != name {
					continue
				}

				checkAppReady(name)

				if cfg.Build.Static.Enable {
					os.MkdirAll(path.Join("dist", name), os.ModePerm)

					cmd := exec.Command("go", "run", path.Join(name, ".mango", "go", "builder.go"))
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						config.Logger.Error().Err(err).Str("app", name).Msg("Build failed")
					}
					continue
				}

				for _, platform := range cfg.Build.Platforms {
					os.MkdirAll(path.Join("dist", name), os.ModePerm)

					filename := filepath.Base(name)

					cmd := exec.Command(
						"env", fmt.Sprintf("GOOS=%s", platform.Os), fmt.Sprintf("GOARCH=%s", platform.Arch), "CGOENABLED=0",
						"go", "build", "-ldflags", "-s -w", "-o", path.Join("dist", name, fmt.Sprintf("%s-%s-%s", filename, platform.Os, platform.Arch)), fmt.Sprintf("./%s", name))
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						config.Logger.Error().Err(err).Str("app", name).Msg("Build failed")
					} else {
						config.Logger.Info().Str("app", name).Str("file", path.Join("dist", name, fmt.Sprintf("%s-%s-%s", filename, platform.Os, platform.Arch))).Msg("Build Succeed")
					}
				}
			}
			config.Logger.Info().Msg("Build Completed !")
		},
	}
	cmd.Flags().StringVarP(&filter, "filter", "f", "", "Only generate certain app")
	return cmd
}
