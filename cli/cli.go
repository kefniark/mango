package main

import (
	"os"

	"github.com/kefniark/mango/cli/config"
	"github.com/kefniark/mango/cli/generate"
	"github.com/kefniark/mango/cli/prepare"
	"github.com/spf13/cobra"
)

var preparer = []config.Executer{}
var generater = []config.Executer{}

func initExec(cfg *config.Config) {
	// preparer
	preparer = append(preparer, prepare.AirPrepare{})
	preparer = append(preparer, prepare.NodeJSPrepare{})
	preparer = append(preparer, prepare.OpenAPIPrepare{Config: cfg})
	preparer = append(preparer, prepare.SQLCPrepare{Config: cfg})
	preparer = append(preparer, prepare.StaticFilePrepare{})

	// generater
	generater = append(generater, generate.ProtoGenerator{})
	generater = append(generater, generate.SQLCGenerator{})
	generater = append(generater, generate.TailwindGenerator{})
	generater = append(generater, generate.TemplGenerator{})
}

func main() {
	cfg := config.Parse()
	initExec(cfg)

	rootCmd := &cobra.Command{
		Use:   "mango",
		Short: "Mango is a Web Development Framework for Go.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}

	rootCmd.AddCommand(buildCmd(cfg))
	rootCmd.AddCommand(cleanCmd(cfg))
	rootCmd.AddCommand(devCmd(cfg))
	rootCmd.AddCommand(prepareCmd(cfg))
	rootCmd.AddCommand(generateCmd(cfg))
	rootCmd.AddCommand(lintCmd(cfg))
	rootCmd.AddCommand(formatCmd(cfg))

	if err := rootCmd.Execute(); err != nil {
		config.Logger.Err(err).Msg("Cannot execute command")
		os.Exit(1)
	}
}
