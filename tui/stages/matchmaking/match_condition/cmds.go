package match_condition

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/dtos"
	matchmakingstages "github.com/yelaco/gchess-tui/tui/stages/matchmaking"
)

func cancelMatchingCondition() tea.Cmd {
	return func() tea.Msg {
		return matchmakingstages.MatchConditionCancelMsg{}
	}
}

func confirmMatchingCondition(condition dtos.MatchCondition) tea.Cmd {
	return func() tea.Msg {
		return matchmakingstages.MatchConditionConfirmMsg{
			Condition: condition,
		}
	}
}
