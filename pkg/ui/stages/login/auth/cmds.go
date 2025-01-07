package auth

import (
	tea "github.com/charmbracelet/bubbletea"
	loginhandler "github.com/yelaco/gchess-tui/pkg/handlers/login"
	loginstages "github.com/yelaco/gchess-tui/pkg/ui/stages/login"
)

func (m AuthStageModel) doLogin() tea.Cmd {
	return func() tea.Msg {
		user, err := loginhandler.LoginUser(m.authInfo)
		if err != nil {
			return loginstages.AuthFailedMsg{Error: err}
		}

		return loginstages.AuthResultMsg{
			User: user,
		}
	}
}

func cancelLogin() tea.Cmd {
	return func() tea.Msg {
		return loginstages.AuthCancelMsg{}
	}
}
