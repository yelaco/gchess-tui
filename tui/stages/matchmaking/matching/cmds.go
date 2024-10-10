package matching

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/dtos"
	"github.com/yelaco/gchess-tui/tui/stages/matchmaking"
	matchmakingstages "github.com/yelaco/gchess-tui/tui/stages/matchmaking"
)

func (m MatchingStageModel) doMatching() tea.Cmd {
	return func() tea.Msg {
		// TODO: implement matching logic
		time.Sleep(5 * time.Second)
		return matchmakingstages.MatchingResultMsg{
			Info: dtos.MatchInfo{},
		}
	}
}

func cancelMatching() tea.Cmd {
	return func() tea.Msg {
		return matchmaking.MatchingCancelMsg{}
	}
}
