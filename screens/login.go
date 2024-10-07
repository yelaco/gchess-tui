package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	loginstages "github.com/yelaco/gchess-tui/stages/login"
)

type LoginScreenModel struct {
	model         tea.Model
	width, height int
}

func NewLoginScreenModel() LoginScreenModel {
	return LoginScreenModel{
		model: loginstages.NewInputStageModel(),
	}
}

func (m LoginScreenModel) Init() tea.Cmd {
	return m.model.Init()
}

func (m LoginScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if Dump != nil {
		fmt.Fprintf(Dump, "LoginScreenModel: %#v\n", msg)
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case loginstages.LoginInfoMsg:
		m.model = loginstages.NewAuthStageModel(msg)
		cmd = tea.Batch(tea.ClearScreen, m.model.Init())
	case loginstages.LoginCancelMsg:
		m.model = loginstages.NewInputStageModel()
		cmd = tea.Batch(tea.ClearScreen, m.model.Init())
	case loginstages.LoginResultMsg:
		return RootScreen().SwitchScreen(NewHomeScreenModel(msg.User))
	default:
		m.model, cmd = m.model.Update(msg)
	}

	return m, cmd
}

func (m LoginScreenModel) View() string {
	header := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		Render("Login")
	footer := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Render("footer")
	content := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height-lipgloss.Height(header)-lipgloss.Height(footer)).
		Align(lipgloss.Center, lipgloss.Center).
		Render(m.model.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
