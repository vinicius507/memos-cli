package memo

import (
	"github.com/spf13/cobra"
	"github.com/vinicius507/memoscli/cmd/memo/create"
)

func NewMemoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "memo",
		Short: "Memo related commands",
	}
	cmd.AddCommand(create.NewCreateCmd())
	return cmd
}
