package gameplay

import (
	tea "github.com/charmbracelet/bubbletea"
	playhandler "github.com/yelaco/gchess-tui/pkg/handlers/play"
	"github.com/yelaco/gchess-tui/pkg/ui/stages/play"
)

func waitForOpponentMove() tea.Cmd {
	return func() tea.Msg {
		err := playhandler.WaitForUpdate()
		if err != nil {
			return nil
		}
		return play.GameUpdatedMsg{}
	}
}

func sendMove(fen, move string) tea.Cmd {
	return func() tea.Msg {
		err := playhandler.SendMove(fen, move)
		if err != nil {
			return nil
		}
		return play.GameUpdatedMsg{}
	}
}
