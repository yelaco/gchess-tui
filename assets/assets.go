package assets

import (
	_ "embed"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
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

var personIcon = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).SetString("👤")

func getPersonIcon() string {
	return personIcon.String()
}

func GetUserFooter(user domains.User) string {
	return fmt.Sprintf("- Logged in as %s -", user.Username)
}
