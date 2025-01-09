package matching

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
	matchmakingstages "github.com/yelaco/gchess-tui/pkg/ui/stages/matchmaking"
)

var cancelHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

type MatchingStageModel struct {
	user           domains.User
	matchCondition domains.MatchCondition
	spinner        spinner.Model
}

func NewMatchingStageModel(matchCondition domains.MatchCondition) tea.Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#4e7837"))

	return MatchingStageModel{
		matchCondition: matchCondition,
		spinner:        s,
	}
}

func (m MatchingStageModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.doMatchmaking())
}

func (m MatchingStageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc", "backspace":
			return m, cancelMatchmaking()
		default:
			return m, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case matchmakingstages.MatchedMsg:
		return m, matchmakingstages.CompleteMatchmaking()
	}

	return m, nil
}

func (m MatchingStageModel) View() string {
	return fmt.Sprintf("\n\n%sMatching...\n\n%s",
		m.spinner.View(),
		cancelHelpStyle.Render("(Press backspace/esc to cancel)"),
	)
}
