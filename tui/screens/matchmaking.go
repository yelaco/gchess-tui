package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/domains/dtos"
	"github.com/yelaco/gchess-tui/tui"
	matchmakingstages "github.com/yelaco/gchess-tui/tui/stages/matchmaking"
	matchconditionstage "github.com/yelaco/gchess-tui/tui/stages/matchmaking/match_condition"
	matchingstage "github.com/yelaco/gchess-tui/tui/stages/matchmaking/matching"
	"github.com/yelaco/gchess-tui/tui/theme"
)

type MatchmakingScreenModel struct {
	user          dtos.User
	stage         tea.Model
	width, height int
}

func NewMatchmakingScreen(user dtos.User) MatchmakingScreenModel {
	m := MatchmakingScreenModel{
		user:  user,
		stage: matchconditionstage.NewMatchingConditionStageModel(user),
	}
	return m
}

func (m MatchmakingScreenModel) Init() tea.Cmd {
	return m.stage.Init()
}

func (m MatchmakingScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	tui.DumpMsgLog("MatchingScreenModel", msg)

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case matchmakingstages.MatchConditionCancelMsg:
		return RootScreen().SwitchScreen(NewHomeScreenModel(m.user))
	case matchmakingstages.MatchConditionConfirmMsg:
		m.stage = matchingstage.NewMatchingStageModel(m.user, msg.Condition)
		return m, tea.Batch(tea.ClearScreen, m.stage.Init())
	case matchmakingstages.MatchingCancelMsg:
		m.stage = matchconditionstage.NewMatchingConditionStageModel(m.user)
		return m, tea.Batch(tea.ClearScreen, m.stage.Init())
	case matchmakingstages.MatchmakingCompleteMsg:
		return RootScreen().SwitchScreen(NewPlayScreenModel(msg.Info))
	default:
		m.stage, cmd = m.stage.Update(msg)
	}

	return m, cmd
}

func (m MatchmakingScreenModel) View() string {
	header := theme.HeaderStyle.Width(m.width).Render("Matching")
	footer := theme.FooterStyle.Width(m.width).Render(m.user.Username)
	content := theme.ContentStyle.
		Width(m.width).
		Height(m.height - lipgloss.Height(header) - lipgloss.Height(footer)).
		Render(m.stage.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
