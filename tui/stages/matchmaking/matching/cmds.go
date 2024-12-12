package matching

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/dtos"
	matchmakingstages "github.com/yelaco/gchess-tui/tui/stages/matchmaking"
)

func (m MatchingStageModel) doMatching() tea.Cmd {
	return func() tea.Msg {
		// TODO: implement matching logic
		time.Sleep(1 * time.Second)
		return matchmakingstages.MatchingResultMsg{
			Info: dtos.MatchInfo{},
		}
	}
}

func cancelMatching() tea.Cmd {
	return func() tea.Msg {
		return matchmakingstages.MatchingCancelMsg{}
	}
}
