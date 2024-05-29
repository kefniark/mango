package main

import (
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
				checkAppPrepared(name)
			}

			for name := range *cfg {
				if filter != "" && filter != name {
					continue
				}

				config.Logger.Debug().Str("app", name).Msg("Code-generation ...")
				for _, exec := range generater {
					config.Logger.Debug().Str("app", name).Str("exec", exec.Name()).Msg("Start")
					if err := exec.Execute(name); err != nil {
						config.Logger.Err(err).Msg("Generate Failed")
					}
					config.Logger.Debug().Str("app", name).Str("exec", exec.Name()).Msg("End")
				}
			}

			config.Logger.Info().Msg("Generate Completed !")
		},
	}
	cmd.Flags().StringVarP(&filter, "filter", "f", "", "Only generate certain app")
	return cmd
}
