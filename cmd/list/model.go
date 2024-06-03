package list

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vinicius507/memos-cli/memos"
	"github.com/vinicius507/memos-cli/ui/styles"
)

type model struct {
	err     error
	client  *memos.Client
	memos   []*memos.Memo
	spinner spinner.Model
	loading bool
}

var _ tea.Model = model{}

func newModel(client *memos.Client) *model {
	sp := spinner.New()
	sp.Spinner = spinner.Points
	sp.Style = styles.LoadingMsg
	return &model{
		client:  client,
		spinner: sp,
		loading: true,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, fetchMemos(m.client))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmdErrorMsg:
		m.loading = false
		m.err = msg.err
		return m, nil
	case memosListMsg:
		m.loading = false
		m.memos = msg.memos
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	var msg string
	if m.err != nil {
		msg = styles.ErrorMsg.Render(m.err.Error()) + "\n"
		msg += "Press q or Ctrl+C to quit."
		return msg
	}
	if m.loading {
		msg += m.spinner.View()
		msg += " Loading memos"
		return msg
	}
	if len(m.memos) == 0 {
		return styles.WarningMsg.Render("No memos found.")
	}
	for _, memo := range m.memos {
		msg += memo.Content + "\n\n"
	}
	return msg
}

type cmdErrorMsg struct{ err error }

type memosListMsg struct{ memos []*memos.Memo }

func fetchMemos(client *memos.Client) tea.Cmd {
	return func() tea.Msg {
		res, err := client.ListMemos()
		if err != nil {
			return cmdErrorMsg{err}
		}
		return memosListMsg{res.Memos}
	}
}
