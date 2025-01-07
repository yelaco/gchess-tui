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
	Info domains.MatchInfo
}

type MatchmakingCompleteMsg struct {
	Info domains.MatchInfo
}

func CompleteMatchmaking(info domains.MatchInfo) tea.Cmd {
	return func() tea.Msg {
		return MatchmakingCompleteMsg{
			Info: info,
		}
	}
}
