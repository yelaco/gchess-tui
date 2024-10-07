package login

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var cancelHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

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
	return tea.Batch(m.spinner.Tick, doLogin(m.authInfo))
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
	case LoginCancelMsg:
		inputStageModel := NewInputStageModel()
		return inputStageModel, tea.Batch(tea.ClearScreen, inputStageModel.Init())
	case LoginResultMsg:
		// return m, completeLogin(msg.user)
		loadInfoStageModel := NewLoadInfoStageModel(msg.user)
		return loadInfoStageModel, tea.Batch(tea.ClearScreen, loadInfoStageModel.Init())
	default:
		return m, nil
	}
}

func (m AuthStageModel) View() string {
	return fmt.Sprintf("\n\n\n%sAuthenticating...\n\n%s",
		m.spinner.View(),
		cancelHelpStyle.Render("(Press ctrl+c/backspace/q to cancel)"),
	)
}
