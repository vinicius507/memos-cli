package memo

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vinicius507/memoscli/cmd/cli"
	"github.com/vinicius507/memoscli/memos"
)

func getEditorCmd() (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}
	if _, err := exec.LookPath("editor"); err != nil {
		return "", err
	}
	return editor, nil
}

func NewCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new memo",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := &memos.MemosClient{
				ServerAddr:  viper.GetString("api_url"),
				AccessToken: viper.GetString("api_token"),
			}
			editorCmd, err := getEditorCmd()
			if err != nil {
				return fmt.Errorf("failed to create memo: %w", err)
			}
			model := cli.NewCreateMemoModel(client, editorCmd)
			_, err = tea.NewProgram(model).Run()
			return err
		},
	}
	return cmd
}
