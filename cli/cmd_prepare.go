package main

import (
	"github.com/kefniark/go-web-server/cli/config"
	"github.com/spf13/cobra"
)

func prepareCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "prepare",
		Short: "Prepare Code for Mango (Generate .mango folder)",
		Run: func(cmd *cobra.Command, args []string) {
			for name := range *cfg {
				config.Logger.Info().Str("app", name).Msg("Prepare .mango folders ...")
				for _, exec := range preparer {
					if err := exec.Execute(name); err != nil {
						config.Logger.Err(err)
						return
					}
				}
			}
		},
	}
}
