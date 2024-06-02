package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vinicius507/memos-cli/ui/styles"
)

func PrintError(err error) tea.Cmd {
	msg := styles.ErrorMsg.Render(err.Error())
	return tea.Println(msg)
}

func PrintWarning(msg string) tea.Cmd {
	msg = styles.WarningMsg.Render(msg)
	return tea.Println(msg)
}

func PrintSuccess(msg string) tea.Cmd {
	msg = styles.SuccessMsg.Render(msg)
	return tea.Println(msg)
}
