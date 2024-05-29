package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/kefniark/mango/cli/config"
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

				// cleanup potential buf files
				entries, _ := os.ReadDir(name)
				for _, e := range entries {
					if strings.HasPrefix(e.Name(), "buf.") {
						os.Remove(path.Join(name, e.Name()))
					}
				}

				// cleanup recursively (templ)
				filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
					if strings.HasSuffix(path, "_templ.go") {
						os.Remove(path)
					}
					return nil
				})
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
