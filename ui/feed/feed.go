package feed

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Items []Item
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	if len(m.Items) == 0 {
		return "Empty"
	}
	str := strings.Builder{}
	for _, item := range m.Items {
		str.WriteString(item.View() + "\n")
	}
	return str.String()
}

type Item struct {
	Content string
}

func (i Item) View() string {
	return i.Content
}
