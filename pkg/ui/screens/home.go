package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/assets"
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
	homestages "github.com/yelaco/gchess-tui/pkg/ui/stages/home"
	menustage "github.com/yelaco/gchess-tui/pkg/ui/stages/home/menu"
	"github.com/yelaco/gchess-tui/pkg/ui/theme"
)

type HomeScreenModel struct {
	user          domains.User
	stage         tea.Model
	width, height int
}

func NewHomeScreenModel(user domains.User) HomeScreenModel {
	return HomeScreenModel{
		user:  user,
		stage: menustage.NewMenuStageModel(),
	}
}

func (m HomeScreenModel) Init() tea.Cmd {
	return m.stage.Init()
}

func (m HomeScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case homestages.PlayMsg:
		return RootScreen().SwitchScreen(NewMatchmakingScreen(m.user))
	case homestages.ViewMatchMsg:
	case homestages.MatchHistoryMsg:
	case homestages.SocialMsg:
	case homestages.LogOutMsg:
		return RootScreen().SwitchScreen(NewLoginScreenModel())
	default:
		m.stage, cmd = m.stage.Update(msg)
	}

	return m, cmd
}

func (m HomeScreenModel) View() string {
	header := theme.HeaderStyle.Width(m.width).Render("Home")
	footer := theme.FooterStyle.Width(m.width).Render(assets.GetUserFooter(m.user))
	content := theme.ContentStyle.
		Width(m.width).
		Height(m.height - lipgloss.Height(header) - lipgloss.Height(footer)).
		Render(m.stage.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
