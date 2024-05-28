package main

import (
	"os"
	"path"
	"strings"

	"github.com/kefniark/go-web-server/cli/config"
	"github.com/spf13/cobra"
)

func cleanCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "clean",
		Short: "Cleanup Mango files (.mango, dist, codegen, ...)",
		Run: func(cmd *cobra.Command, args []string) {
			config.Logger.Info().Msg("Cleanup Folders ...")
			for name := range *cfg {
				os.RemoveAll(path.Join(name, ".mango"))
				os.RemoveAll(path.Join(name, "codegen"))
			}
			os.RemoveAll(path.Join(".", "dist"))

			// Remove Local SQlite DB
			for _, cfg := range *cfg {
				if !strings.HasSuffix(cfg.Database.URL, ".db") {
					continue
				}
				os.Remove(path.Join(".", cfg.Database.URL))
			}

			// Remove Nodejs
			for name := range *cfg {
				os.RemoveAll(path.Join(name, "tools", "nodejs", "node_modules"))
			}

			config.Logger.Info().Msg("Clean Completed !")
		},
	}
}
