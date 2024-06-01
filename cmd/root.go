package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vinicius507/memos-cli/cmd/create"
)

func Execute() {
	cmd := &cobra.Command{
		Use:   "memos",
		Short: "A CLI client for Memos",
	}
	cmd.AddCommand(create.New())
	if err := cmd.Execute(); err != nil {
		cmd.PrintErr(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
