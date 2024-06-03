package create

import (
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vinicius507/memos-cli/memos"
	"github.com/vinicius507/memos-cli/ui/cmd"
	"github.com/vinicius507/memos-cli/ui/styles"
)

type cmdErrorMsg struct{ err error }

func (e cmdErrorMsg) Error() string {
	return e.err.Error()
}

type model struct {
	client    *memos.Client
	editorCmd string
	spinner   spinner.Model
	saving    bool
}

var _ tea.Model = model{}

func newModel(client *memos.Client, editorCmd string) model {
	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = styles.LoadingMsg
	return model{client: client, editorCmd: editorCmd, spinner: sp}
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

func saveMemo(client *memos.Client, file string) tea.Cmd {
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
			cmd.PrintError(msg.err),
			tea.Quit,
		)
	case tempFileMsg:
		return m, openEditor(m.editorCmd, msg.file)
	case editorFinishedMsg:
		m.saving = true
		return m, saveMemo(m.client, msg.file)
	case memoIsEmptyMsg:
		return m, tea.Sequence(
			cmd.PrintWarning("Nothing to save. No memo was created."),
			tea.Quit,
		)
	case memoSavedMsg:
		return m, tea.Sequence(
			cmd.PrintSuccess("Memo created successfully!"),
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
	if m.saving {
		return m.spinner.View() + " Saving memo..."
	}
	return ""
}
