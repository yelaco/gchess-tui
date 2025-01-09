package matching

import (
	tea "github.com/charmbracelet/bubbletea"
	matchmakinghandlers "github.com/yelaco/gchess-tui/pkg/handlers/matchmaking"
	matchmakingstages "github.com/yelaco/gchess-tui/pkg/ui/stages/matchmaking"
)

func (m MatchingStageModel) doMatchmaking() tea.Cmd {
	return func() tea.Msg {
		err := matchmakinghandlers.Matchmaking()
		if err != nil {
			return nil
		}
		return matchmakingstages.MatchedMsg{}
	}
}

func cancelMatchmaking() tea.Cmd {
	return func() tea.Msg {
		return matchmakingstages.MatchingCancelMsg{}
	}
}
