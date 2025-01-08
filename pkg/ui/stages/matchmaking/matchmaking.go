package matchmaking

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
)

type MatchConditionConfirmMsg struct {
	Condition domains.MatchCondition
}

type MatchConditionCancelMsg struct{}

type MatchingCancelMsg struct{}

type MatchingResultMsg struct {
	Info domains.Match
}

type MatchmakingCompleteMsg struct {
	Info domains.Match
}

func CompleteMatchmaking(info domains.Match) tea.Cmd {
	return func() tea.Msg {
		return MatchmakingCompleteMsg{
			Info: info,
		}
	}
}
