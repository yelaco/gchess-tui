package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	loginstages "github.com/yelaco/gchess-tui/pkg/ui/stages/login"
	formstage "github.com/yelaco/gchess-tui/pkg/ui/stages/login/form"
	"github.com/yelaco/gchess-tui/pkg/ui/theme"
)

type LoginScreenModel struct {
	stage         tea.Model
	width, height int
}

func NewLoginScreenModel() LoginScreenModel {
	return LoginScreenModel{
		stage: formstage.NewFormStageModel(),
	}
}

func (m LoginScreenModel) Init() tea.Cmd {
	return m.stage.Init()
}

func (m LoginScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case loginstages.AuthCancelMsg:
		m.stage = formstage.NewFormStageModel()
		return m, tea.Batch(tea.ClearScreen, m.stage.Init())
	case loginstages.AuthFailedMsg:
		return m, tea.Quit
	case loginstages.LoginCompleteMsg:
		return RootScreen().SwitchScreen(NewHomeScreenModel(msg.User))
	default:
		m.stage, cmd = m.stage.Update(msg)
	}

	return m, cmd
}

func (m LoginScreenModel) View() string {
	header := theme.HeaderStyle.Width(m.width).Render("Login")
	footer := theme.FooterStyle.Width(m.width).Render("Welcome to gchess!")
	content := theme.ContentStyle.
		Width(m.width).
		Height(m.height - lipgloss.Height(header) - lipgloss.Height(footer)).
		Render(m.stage.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
