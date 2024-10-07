package login

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/user"
)

// Submit info
type LoginInfoMsg struct {
	username string
	password string
}

func submitLoginInfo(info LoginInfoMsg) tea.Cmd {
	return func() tea.Msg {
		return info
	}
}

// Authenticate
type LoginResultMsg struct {
	User user.User
}

func authenticate(info LoginInfoMsg) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(3 * time.Second)
		return LoginResultMsg{
			User: user.User{
				Username:  info.username,
				Email:     "test@gmail.com",
				Rating:    100,
				CreatedAt: time.Now(),
			},
		}
	}
}

// Cancel
type LoginCancelMsg struct{}

func cancelLogin() tea.Cmd {
	return func() tea.Msg {
		return LoginCancelMsg{}
	}
}
