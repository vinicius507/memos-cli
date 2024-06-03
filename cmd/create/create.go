package create

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/vinicius507/memos-cli/config"
	"github.com/vinicius507/memos-cli/memos"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create a new memo",
		Aliases: []string{"c", "new"},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.GetConfig()
			client := &memos.Client{
				ServerAddr:  cfg.Api.Url,
				AccessToken: cfg.Api.AccessToken,
			}
			editorCmd, err := getEditorCmd()
			if err != nil {
				return fmt.Errorf("failed to create memo: %w", err)
			}
			model := newModel(client, editorCmd)
			if _, err = tea.NewProgram(model).Run(); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func getEditorCmd() (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}
	if _, err := exec.LookPath(editor); err != nil {
		return "", err
	}
	return editor, nil
}
