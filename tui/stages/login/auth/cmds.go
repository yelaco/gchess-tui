package auth

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/dtos"
	"github.com/yelaco/gchess-tui/tui"
	loginstages "github.com/yelaco/gchess-tui/tui/stages/login"
)

func (m AuthStageModel) doLogin() tea.Cmd {
	return func() tea.Msg {
		user, err := tui.GetApp().LoginUsecase.LoginUser(m.authInfo)
		if err != nil {
			return loginstages.AuthFailedMsg{Error: err}
		}

		return loginstages.AuthResultMsg{
			User: dtos.UserEntityToDto(user),
		}
	}
}

func cancelLogin() tea.Cmd {
	return func() tea.Msg {
		return loginstages.AuthCancelMsg{}
	}
}
