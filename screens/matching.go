package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/domains/user"
	matchingconditionstage "github.com/yelaco/gchess-tui/stages/matching/matching_condition"
	"github.com/yelaco/gchess-tui/theme"
)

type MatchingScreenModel struct {
	user          user.User
	stage         tea.Model
	width, height int
}

func NewMatchingScreenModel(user user.User) MatchingScreenModel {
	m := MatchingScreenModel{
		user:  user,
		stage: matchingconditionstage.NewMatchingConditionStageModel(user),
	}
	return m
}

func (m MatchingScreenModel) Init() tea.Cmd {
	return m.stage.Init()
}

func (m MatchingScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if Dump != nil {
		fmt.Fprintf(Dump, "MatchingScreenModel: %#v\n", msg)
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	default:
		m.stage, cmd = m.stage.Update(msg)
	}

	return m, cmd
}

func (m MatchingScreenModel) View() string {
	header := theme.HeaderStyle.Width(m.width).Render("Matching")
	footer := theme.FooterStyle.Width(m.width).Render(m.user.Username)
	content := theme.ContentStyle.
		Width(m.width).
		Height(m.height - lipgloss.Height(header) - lipgloss.Height(footer)).
		Render(m.stage.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
