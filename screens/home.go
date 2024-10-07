package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/domains/user"
	homestages "github.com/yelaco/gchess-tui/stages/home"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

type HomeScreenModel struct {
	user          user.User
	stage         tea.Model
	width, height int
}

func NewHomeScreenModel(user user.User) HomeScreenModel {
	return HomeScreenModel{
		user:  user,
		stage: homestages.NewMenuStageModel(),
	}
}

func (m HomeScreenModel) Init() tea.Cmd {
	return m.stage.Init()
}

func (m HomeScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if Dump != nil {
		fmt.Fprintf(Dump, "HomeScreenModel: %#v\n", msg)
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case homestages.PlayMsg:
		return RootScreen().SwitchScreen(NewMatchingScreenModel(
			m.user,
		))
	case homestages.ViewMatchMsg:
	case homestages.MatchHistoryMsg:
	case homestages.SocialMsg:
	case homestages.LogOutMsg:
	default:
		m.stage, cmd = m.stage.Update(msg)
	}

	return m, cmd
}

func (m HomeScreenModel) View() string {
	header := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		Render("Home")
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
