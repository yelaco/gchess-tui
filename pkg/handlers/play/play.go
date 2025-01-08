package play

import "github.com/yelaco/gchess-tui/pkg/app"

func SendMove(move string) error {
	app.GetMatch().MoveCh <- move
	// TODO: wait for response
	return nil
}
