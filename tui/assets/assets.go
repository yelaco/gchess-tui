package assets

import (
	_ "embed"

	"github.com/charmbracelet/lipgloss"
)

//go:embed logo.txt
var logo string

func GetLogo() string {
	return logo
}

var checkMark = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")

func GetCheckMark() string {
	return checkMark.String()
}

var xMark = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).SetString("✗")

func GetXMark() string {
	return xMark.String()
}
