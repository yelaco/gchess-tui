package login

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AuthStageModel struct {
	authInfo LoginInfoMsg
	spinner  spinner.Model
}

func NewAuthStageModel(info LoginInfoMsg) AuthStageModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return AuthStageModel{
		authInfo: info,
		spinner:  s,
	}
}

func (m AuthStageModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, authenticate(m.authInfo))
}

func (m AuthStageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "backspace", "q":
			return m, cancelLogin()
		default:
			return m, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m AuthStageModel) View() string {
	return fmt.Sprintf("%sAuthenticating...press ctrl+c to cancel\n\n", m.spinner.View())
}
