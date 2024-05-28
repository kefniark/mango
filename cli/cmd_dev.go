package main

import (
	"os"
	"os/exec"
	"path"

	"github.com/kefniark/go-web-server/cli/config"
	"github.com/spf13/cobra"
)

func devCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "dev",
		Short: "Start Dev Servers (with hot reload)",
		Run: func(cmd *cobra.Command, args []string) {
			config.Logger.Info().Msg("Start Dev Servers ...")
			cmds := []*exec.Cmd{}
			for name := range *cfg {
				if _, err := os.Stat(path.Join(name, ".mango")); err != nil {
					config.Logger.Warn().Msg("Need preparation first, wait a moment ...")
					for _, exec := range preparer {
						exec.Execute(name)
					}
					for _, exec := range generater {
						exec.Execute(name)
					}
				}

				cmd := exec.Command("air", "-c", path.Join(name, ".mango", "air.toml"))
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmds = append(cmds, cmd)
				cmd.Start()
			}

			for _, cmd := range cmds {
				if err := cmd.Wait(); err != nil {
					config.Logger.Error().Err(err).Msg("Dev Servers failed")
				}
			}
		},
	}
}
