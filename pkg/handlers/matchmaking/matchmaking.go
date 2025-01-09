package handlers

import (
	"github.com/yelaco/gchess-tui/pkg/app"
	"github.com/yelaco/gchess-tui/pkg/gchessclient/wsclient"
)

func Matchmaking() error {
	// TODO: Set client based on match type
	client := wsclient.NewWsClient()
	app.SetClient(client)

	err := client.Matchmaking()
	if err != nil {
		return err
	}
	return nil
}
