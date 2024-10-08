package auth

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/dtos"
	loginstages "github.com/yelaco/gchess-tui/tui/stages/login"
)

func doLogin(info dtos.Login) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(3 * time.Second)
		return loginstages.AuthResultMsg{
			User: dtos.User{
				Username:  info.Username,
				Email:     "test@gmail.com",
				Rating:    100,
				CreatedAt: time.Now(),
			},
		}
	}
}

func cancelLogin() tea.Cmd {
	return func() tea.Msg {
		return loginstages.AuthCancelMsg{}
	}
}
