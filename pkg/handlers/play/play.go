package play

import "github.com/yelaco/gchess-tui/pkg/app"

func SendMove(fen, move string) error {
	app.GetMatch().MoveCh <- fen + " " + move
	return <-app.GetMatch().ErrorCh
}
