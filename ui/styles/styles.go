package styles

import "github.com/charmbracelet/lipgloss"

var (
	errorIcon   = icon("", lipgloss.Color("9"))
	warningIcon = icon("", lipgloss.Color("11"))
	successIcon = icon("󰄬", lipgloss.Color("10"))

	ErrorMsg   = msgStyle(errorIcon)
	WarningMsg = msgStyle(warningIcon)
	SuccessMsg = msgStyle(successIcon)

	LoadingMsg = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
)

func icon(icon string, color lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(color).SetString(icon)
}

func msgStyle(iconStyle lipgloss.Style) lipgloss.Style {
	return lipgloss.NewStyle().SetString(iconStyle.Render())
}
