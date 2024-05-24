package cli

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vinicius507/memoscli/memos"
)

type CreateMemoModel struct {
	client        *memos.MemosClient
	editorCommand string
}

var _ tea.Model = CreateMemoModel{}

func NewCreateMemoModel(client *memos.MemosClient, editorCommand string) CreateMemoModel {
	return CreateMemoModel{client, editorCommand}
}

type tempFileMsg struct{ file string }

func createTempFile() tea.Msg {
	file, err := os.CreateTemp("", "memos-cli-*.md")
	if err != nil {
		return cmdErrorMsg{err}
	}
	defer file.Close()
	return tempFileMsg{file.Name()}
}

type editorFinishedMsg struct{ file string }

func openEditor(editor, file string) tea.Cmd {
	cmd := exec.Command(editor, file)
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		if err != nil {
			return cmdErrorMsg{err}
		}
		return editorFinishedMsg{file}
	})
}

type memoSavedMsg struct{ memo *memos.Memo }

func saveMemo(client *memos.MemosClient, file string) tea.Cmd {
	return func() tea.Msg {
		content, err := os.ReadFile(file)
		if err != nil {
			return cmdErrorMsg{err: fmt.Errorf("could not create memo: %w", err)}
		}
		if len(content) == 0 {
			return tea.Quit() // If the file is empty, just quit
		}
		memo, err := client.NewMemo(string(content))
		if err != nil {
			return cmdErrorMsg{err}
		}
		return memoSavedMsg{memo}
	}
}

func (m CreateMemoModel) Init() tea.Cmd {
	return createTempFile
}

func (m CreateMemoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmdErrorMsg:
		fmt.Printf("Error: could not create memo: %v\n", msg.err)
		return m, tea.Quit
	case tempFileMsg:
		return m, openEditor(m.editorCommand, msg.file)
	case editorFinishedMsg:
		return m, saveMemo(m.client, msg.file)
	case memoSavedMsg:
		return m, tea.Quit
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m CreateMemoModel) View() string {
	return ""
}
