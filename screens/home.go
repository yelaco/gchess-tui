package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/domains/user"
	homestages "github.com/yelaco/gchess-tui/stages/home"
	"github.com/yelaco/gchess-tui/theme"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#4e7837"))
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
	header := theme.HeaderStyle.Width(m.width).Render("Home")
	footer := theme.FooterStyle.Width(m.width).Render(m.user.Username)
	content := theme.ContentStyle.
		Width(m.width).
		Height(m.height - lipgloss.Height(header) - lipgloss.Height(footer)).
		Render(m.stage.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
