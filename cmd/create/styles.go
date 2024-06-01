package create

import "github.com/charmbracelet/lipgloss"

var (
	errorIcon   = icon("", lipgloss.Color("9"))
	warningIcon = icon("", lipgloss.Color("11"))
	successIcon = icon("󰄬", lipgloss.Color("10"))

	errorStyle   = lipgloss.NewStyle().SetString(errorIcon.Render())
	warningStyle = lipgloss.NewStyle().SetString(warningIcon.Render())
	successStyle = lipgloss.NewStyle().SetString(successIcon.Render())
)

func icon(icon string, color lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(color).SetString(icon)
}
