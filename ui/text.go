package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).SetString("").MarginRight(1)
	warningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("11")).SetString("").MarginRight(1)
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).SetString("󰄬").MarginRight(1)
)

func RenderError(err error) string {
	return lipgloss.JoinHorizontal(lipgloss.Bottom, errorStyle.Render(), err.Error())
}

func RenderWarning(msg string) string {
	return lipgloss.JoinHorizontal(lipgloss.Bottom, warningStyle.Render(), msg)
}

func RenderSuccess(msg string) string {
	return lipgloss.JoinHorizontal(lipgloss.Bottom, successStyle.Render(), msg)
}
