package feed

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Items    []Item
	viewport viewport.Model
	ready    bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height)
			m.viewport.HighPerformanceRendering = true
			m.viewport.SetContent("Empty")
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height
		}
	case setItemsMsg:
		m.Items = msg.Items
		m.updateViewPortContent()
		cmds = append(cmds, viewport.Sync(m.viewport))
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.viewport.View()
}

func (m *Model) updateViewPortContent() {
	str := strings.Builder{}
	for _, item := range m.Items {
		str.WriteString(item.View() + "\n")
	}
	m.viewport.SetContent(str.String())
}

type setItemsMsg struct {
	Items []Item
}

func (m Model) SetItems(items []Item) tea.Cmd {
	return func() tea.Msg {
		return setItemsMsg{items}
	}
}

type Item struct {
	Content string
}

func (i Item) View() string {
	return i.Content
}
