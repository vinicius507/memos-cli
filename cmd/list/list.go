package list

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/vinicius507/memos-cli/config"
	"github.com/vinicius507/memos-cli/memos"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List your memos",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.GetConfig()
			client := &memos.Client{
				ServerAddr:  cfg.Api.Url,
				AccessToken: cfg.Api.AccessToken,
			}
			model := newModel(client)
			if _, err := tea.NewProgram(model).Run(); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
