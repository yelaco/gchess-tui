package theme

import "github.com/charmbracelet/lipgloss"

var HeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#4e7837")).
	Align(lipgloss.Center).
	Border(lipgloss.HiddenBorder(), false, false, true, false).
	Bold(true)

var FooterStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("249")).
	Align(lipgloss.Center)

var ContentStyle = lipgloss.NewStyle().
	Align(lipgloss.Center, lipgloss.Center)
