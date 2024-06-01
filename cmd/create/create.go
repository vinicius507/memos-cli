package create

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vinicius507/memos-cli/memos"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create a new memo",
		Aliases: []string{"c", "new"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !viper.IsSet("api.url") {
				return fmt.Errorf("missing api_url in config file")
			}
			if !viper.IsSet("api.token") {
				return fmt.Errorf("missing api_token in config file")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client := &memos.MemosClient{
				ServerAddr:  viper.GetString("api.url"),
				AccessToken: viper.GetString("api.token"),
			}
			editorCmd, err := getEditorCmd()
			if err != nil {
				return fmt.Errorf("failed to create memo: %w", err)
			}
			model := newModel(client, editorCmd)
			_, err = tea.NewProgram(model).Run()
			return err
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
