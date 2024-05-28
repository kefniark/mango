package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/kefniark/mango/cli/config"
	"github.com/spf13/cobra"
)

func formatCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "format",
		Short: "Auto-format code",
		Run: func(cmd *cobra.Command, args []string) {
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

				config.Logger.Info().Str("app", name).Msg("Lint code with Auto-fix ...")
				if err := lintFormat(name); err != nil {
					config.Logger.Err(err)
					return
				}
			}

			if err := prettier(true); err != nil {
				config.Logger.Err(err)
				return
			}

			config.Logger.Info().Msg("Format Completed !")
		},
	}
}

func lintCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "lint",
		Short: "Lint code",
		Run: func(cmd *cobra.Command, args []string) {
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

				config.Logger.Info().Str("app", name).Msg("Lint code ...")
				if err := lint(name); err != nil {
					config.Logger.Err(err)
					return
				}
			}

			if err := prettier(false); err != nil {
				config.Logger.Err(err)
				return
			}

			config.Logger.Info().Msg("Lint Completed !")
		},
	}
}

func prettier(autofix bool) error {
	arg := "--check"
	if autofix {
		arg = "--write"
	}
	golangciCmd := exec.Command("prettier", "**/*.{js,css,json,yaml,md}", arg)
	golangciCmd.Stdout = os.Stdout
	golangciCmd.Stderr = os.Stderr

	return golangciCmd.Run()
}

func lint(name string) error {
	dir, err := filepath.Abs(name)
	if err != nil {
		return err
	}

	golangciCmd := exec.Command("golangci-lint", "run", fmt.Sprintf("%s/...", dir), "--config", fmt.Sprintf("%s/.mango/go/golangci.yaml", dir))
	golangciCmd.Dir = dir
	golangciCmd.Stdout = os.Stdout
	golangciCmd.Stderr = os.Stderr

	return golangciCmd.Run()
}

func lintFormat(name string) error {
	dir, err := filepath.Abs(name)
	if err != nil {
		return err
	}

	golangciCmd := exec.Command("golangci-lint", "run", fmt.Sprintf("%s/...", dir), "--fix", "--config", fmt.Sprintf("%s/.mango/go/golangci.yaml", dir))
	golangciCmd.Stdout = os.Stdout
	golangciCmd.Stderr = os.Stderr
	golangciCmd.Dir = dir
	if err = golangciCmd.Run(); err != nil {
		return err
	}

	sqlcCmd := exec.Command("sqlc", "vet", "-f", fmt.Sprintf("%s/.mango/db/sqlc.yaml", dir))
	golangciCmd.Stdout = os.Stdout
	golangciCmd.Stderr = os.Stderr
	sqlcCmd.Dir = dir
	return sqlcCmd.Run()
}
