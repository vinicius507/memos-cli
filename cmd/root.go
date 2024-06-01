package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vinicius507/memos-cli/cmd/create"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "memos",
		Short: "A CLI client for Memos",
	}
	cmd.AddCommand(create.New())
	return cmd
}

func init() {
	cobra.OnInitialize(initConfig)
}
