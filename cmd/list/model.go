package list

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vinicius507/memos-cli/memos"
	"github.com/vinicius507/memos-cli/ui/feed"
	"github.com/vinicius507/memos-cli/ui/styles"
)

type model struct {
	err     error
	client  *memos.Client
	memos   feed.Model
	spinner spinner.Model
	loading bool
}

var _ tea.Model = model{}

func newModel(client *memos.Client) *model {
	sp := spinner.New()
	sp.Spinner = spinner.Points
	sp.Style = styles.LoadingMsg
	memos := feed.Model{}

	return &model{
		client:  client,
		loading: true,
		memos:   memos,
		spinner: sp,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, fetchMemos(m.client))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case cmdErrorMsg:
		m.loading = false
		m.err = msg.err
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case memosListMsg:
		m.loading = false
		for _, memo := range msg.memos {
			m.memos.Items = append(m.memos.Items, feed.Item{Content: memo.Content})
		}
		return m, nil
	}

	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)
	m.memos, cmd = m.memos.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf(
			"%s\nPress q or Ctrl+C to quit.",
			styles.ErrorMsg.Render(m.err.Error())+"\n",
		)
	}
	if m.loading {
		return m.spinner.View() + " Loading memos"
	}
	return m.memos.View()
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
