package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/domains/dtos"
	"github.com/yelaco/gchess-tui/tui"
	homestages "github.com/yelaco/gchess-tui/tui/stages/home"
	"github.com/yelaco/gchess-tui/tui/theme"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#4e7837"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

type HomeScreenModel struct {
	user          dtos.User
	stage         tea.Model
	width, height int
}

func NewHomeScreenModel(user dtos.User) HomeScreenModel {
	return HomeScreenModel{
		user:  user,
		stage: homestages.NewMenuStageModel(),
	}
}

func (m HomeScreenModel) Init() tea.Cmd {
	return m.stage.Init()
}

func (m HomeScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	tui.DumpMsgLog("HomeScreenModel", msg)

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
		return RootScreen().SwitchScreen(NewLoginScreenModel())
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
