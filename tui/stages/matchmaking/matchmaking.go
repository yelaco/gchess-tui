package matchmaking

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/dtos"
)

type MatchingType string

var (
	MatchingOpponent MatchingType = "Matching Opponent"
	MatchingAi       MatchingType = "Matching AI"
	MatchingFriend   MatchingType = "Matching Friend"
)

type MatchConditionConfirmMsg struct {
	Condition dtos.MatchCondition
}

type MatchConditionCancelMsg struct{}

type MatchingCancelMsg struct{}

type MatchingResultMsg struct {
	Info dtos.MatchInfo
}

type MatchmakingCompleteMsg struct {
	Info dtos.MatchInfo
}

func CompleteMatchmaking(info dtos.MatchInfo) tea.Cmd {
	return func() tea.Msg {
		return MatchmakingCompleteMsg{
			Info: info,
		}
	}
}
