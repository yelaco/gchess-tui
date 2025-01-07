package gameplay

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/pkg/ui/stages/play"
)

func sendMove(fen, move string) tea.Cmd {
	return func() tea.Msg {
		return play.GameUpdateMsg{
			Accepted:    true,
			BoardFen:    "4k2r/6r1/8/8/8/8/3R4/R3K3",
			IsWhiteTurn: true,
		}
	}
}
