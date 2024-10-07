package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/domains/user"
	matchingconditionstage "github.com/yelaco/gchess-tui/stages/matching/matching_condition"
)

type MatchingScreenModel struct {
	user          user.User
	model         tea.Model
	width, height int
}

func NewMatchingScreenModel(user user.User) MatchingScreenModel {
	m := MatchingScreenModel{
		user:  user,
		model: matchingconditionstage.NewMatchingConditionStageModel(user),
	}
	return m
}

func (m MatchingScreenModel) Init() tea.Cmd {
	return m.model.Init()
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
		m.model, cmd = m.model.Update(msg)
	}

	return m, cmd
}

func (m MatchingScreenModel) View() string {
	header := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		Render("Matching")
	footer := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Render("footer")
	content := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height-lipgloss.Height(header)-lipgloss.Height(footer)).
		Align(lipgloss.Center, lipgloss.Top).
		Render(m.model.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
