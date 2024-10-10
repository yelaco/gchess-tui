package auth

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/domains/dtos"
	loginstages "github.com/yelaco/gchess-tui/tui/stages/login"
	loadinfostage "github.com/yelaco/gchess-tui/tui/stages/login/load_info"
)

var cancelHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

type AuthStageModel struct {
	authInfo dtos.Login
	spinner  spinner.Model
}

func NewAuthStageModel(info dtos.Login) AuthStageModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#4e7837"))
	return AuthStageModel{
		authInfo: info,
		spinner:  s,
	}
}

func (m AuthStageModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.doLogin())
}

func (m AuthStageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc", "backspace":
			return m, cancelLogin()
		default:
			return m, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case loginstages.AuthResultMsg:
		// TODO: validate returned user info
		// if msg.User.Username != m.authInfo.Username {}
		loadInfoStageModel := loadinfostage.NewLoadInfoStageModel(msg.User)
		return loadInfoStageModel, tea.Batch(tea.ClearScreen, loadInfoStageModel.Init())
	default:
		return m, nil
	}
}

func (m AuthStageModel) View() string {
	return fmt.Sprintf("\n\n%sAuthenticating...\n\n%s",
		m.spinner.View(),
		cancelHelpStyle.Render("(Press backspace/esc to cancel)"),
	)
}
