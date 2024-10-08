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
	header := lipgloss.NewStyle().
		Background(lipgloss.Color("#4e7837")).
		Foreground(lipgloss.Color("255")).
		Align(lipgloss.Center).
		Width(m.width).
		Bold(true).
		Render("Matching")
	footer := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Render("footer")
	content := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height-lipgloss.Height(header)-lipgloss.Height(footer)).
		Align(lipgloss.Center, lipgloss.Top).
		Render(m.stage.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
