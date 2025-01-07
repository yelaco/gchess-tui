package login

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
)

// Authentication stage
type AuthResultMsg struct {
	User domains.User
}

type AuthCancelMsg struct{}

type AuthFailedMsg struct {
	Error error
}

// Load info stage
type LoadInfoFinishedMsg struct {
	Name     string
	Duration time.Duration
	Result   interface{}
}

type LoadInfoFailedMsg struct {
	Name     string
	Duration time.Duration
}

// Complete screen
type LoginCompleteMsg struct {
	User domains.User
}

func CompleteLogin(user domains.User) tea.Cmd {
	return func() tea.Msg {
		return LoginCompleteMsg{
			User: user,
		}
	}
}
