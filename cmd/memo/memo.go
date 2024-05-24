package memo

import (
	"github.com/spf13/cobra"
)

func NewMemoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "memo",
		Short: "Memo related commands",
	}
	cmd.AddCommand(NewCreateCmd())
	return cmd
}
