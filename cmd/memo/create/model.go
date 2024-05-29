package create

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vinicius507/memos-cli/memos"
)

type cmdErrorMsg struct{ err error }

func (e cmdErrorMsg) Error() string {
	return e.err.Error()
}

type model struct {
	client    *memos.MemosClient
	editorCmd string
}

var _ tea.Model = model{}

func newModel(client *memos.MemosClient, editorCmd string) model {
	return model{client: client, editorCmd: editorCmd}
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

func openEditor(editorCmd, file string) tea.Cmd {
	execEditor := exec.Command(editorCmd, file)
	return tea.ExecProcess(execEditor, func(err error) tea.Msg {
		if err != nil {
			return cmdErrorMsg{err}
		}
		return editorFinishedMsg{file}
	})
}

type memoIsEmptyMsg struct{}

type memoSavedMsg struct{ memo *memos.Memo }

func saveMemo(client *memos.MemosClient, file string) tea.Cmd {
	return func() tea.Msg {
		content, err := os.ReadFile(file)
		if err != nil {
			return cmdErrorMsg{err}
		}
		if len(content) == 0 {
			return memoIsEmptyMsg{}
		}
		memo, err := client.NewMemo(string(content))
		if err != nil {
			return cmdErrorMsg{err}
		}
		return memoSavedMsg{memo}
	}
}

func (m model) Init() tea.Cmd {
	return createTempFile
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmdErrorMsg:
		return m, tea.Sequence(
			tea.Println(errorStyle.Render(msg.Error())),
			tea.Quit,
		)
	case tempFileMsg:
		return m, openEditor(m.editorCmd, msg.file)
	case editorFinishedMsg:
		return m, saveMemo(m.client, msg.file)
	case memoIsEmptyMsg:
		return m, tea.Sequence(
			tea.Println(warningStyle.Render("Nothing to save. No memo was created.")),
			tea.Quit,
		)
	case memoSavedMsg:
		return m, tea.Sequence(
			tea.Println(successStyle.Render("Memo created successfully!")),
			tea.Quit,
		)
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return ""
}
