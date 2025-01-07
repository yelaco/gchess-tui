package match_condition

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
	matchmakingstages "github.com/yelaco/gchess-tui/pkg/ui/stages/matchmaking"
)

func cancelMatchingCondition() tea.Cmd {
	return func() tea.Msg {
		return matchmakingstages.MatchConditionCancelMsg{}
	}
}

func confirmMatchingCondition(condition domains.MatchCondition) tea.Cmd {
	return func() tea.Msg {
		return matchmakingstages.MatchConditionConfirmMsg{
			Condition: condition,
		}
	}
}
