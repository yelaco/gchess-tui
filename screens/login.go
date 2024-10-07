package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	loginstages "github.com/yelaco/gchess-tui/stages/login"
)

type LoginScreenModel struct {
	stage         tea.Model
	width, height int
}

func NewLoginScreenModel() LoginScreenModel {
	return LoginScreenModel{
		stage: loginstages.NewInputStageModel(),
	}
}

func (m LoginScreenModel) Init() tea.Cmd {
	return m.stage.Init()
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
	case loginstages.LoginCompleteMsg:
		return RootScreen().SwitchScreen(NewHomeScreenModel(msg.User))
	default:
		m.stage, cmd = m.stage.Update(msg)
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
		Render(m.stage.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
