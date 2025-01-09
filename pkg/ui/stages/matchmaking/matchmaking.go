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

type MatchedMsg struct{}

type MatchmakingCompleteMsg struct{}

func CompleteMatchmaking() tea.Cmd {
	return func() tea.Msg {
		return MatchmakingCompleteMsg{}
	}
}
