package main

import (
	"os"
	"path"

	"github.com/kefniark/mango/cli/config"
	"github.com/spf13/cobra"
)

func generateCmd(cfg *config.Config) *cobra.Command {
	var filter string
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Code Generation for Mango (Protobuf, SQLC, Templ, ...)",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if prepare is required
			for name := range *cfg {
				if filter != "" && filter != name {
					continue
				}
				if _, err := os.Stat(path.Join(name, ".mango")); err != nil {
					config.Logger.Warn().Msg("Need preparation first, wait a moment ...")
					for _, exec := range preparer {
						exec.Execute(name)
					}
				}
			}

			for name := range *cfg {
				if filter != "" && filter != name {
					continue
				}

				config.Logger.Debug().Str("app", name).Msg("Code-generation ...")
				for _, exec := range generater {
					exec.Execute(name)
				}
			}

			config.Logger.Info().Msg("Generate Completed !")
		},
	}
	cmd.Flags().StringVarP(&filter, "filter", "f", "", "Only generate certain app")
	return cmd
}
