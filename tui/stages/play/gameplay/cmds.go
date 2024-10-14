package gameplay

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/tui/stages/play"
)

func sendMove(fen, move string) tea.Cmd {
	return func() tea.Msg {
		return play.GameUpdateMsg{
			Accepted:    true,
			Fen:         fen,
			IsWhiteTurn: true,
		}
	}
}
