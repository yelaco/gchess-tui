package screens

import (
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
)

func NewPlayScreenDebug() PlayScreenModel {
	return NewPlayScreenModel(domains.Match{
		GameState: domains.GameState{
			Status:      "ACTIVE",
			BoardFen:    "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			IsWhiteTurn: true,
		},
	})
}
