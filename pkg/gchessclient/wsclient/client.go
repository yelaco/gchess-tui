package wsclient

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yelaco/gchess-tui/pkg/app"
	"github.com/yelaco/gchess-tui/pkg/gchessclient"
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
	"github.com/yelaco/gchess-tui/pkg/utils"
)

type client struct {
	conn *websocket.Conn
}

func NewWsClient() gchessclient.Client {
	return &client{}
}

// connect function    attempts to connect to websocket server with retry logic.
func connect(retry int, waitTime time.Duration) (*websocket.Conn, error) {
	wsUrl := app.GetConfig().ServiceWsUrl.String()
	for range retry {
		conn, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
		if err != nil {
			<-time.After(waitTime)
			continue
		}
		return conn, err
	}
	return nil, ErrWsConnectFailed
}

// Matchmaking method    implements matchmaking logic for websocket connection
func (c *client) Matchmaking() error {
	conn, err := connect(5, 5*time.Second)
	if err != nil {
		return fmt.Errorf("%w: %w", gchessclient.ErrMatchmakingFailed, err)
	}
	c.conn = conn

	playerId := utils.GeneratePlayerId()
	err = c.conn.WriteJSON(matchmakingRequest{
		Action: "matching",
		Data: matchmakingData{
			PlayerId: playerId,
		},
	})
	if err != nil {
		return fmt.Errorf("%w: %w", gchessclient.ErrMatchmakingFailed, err)
	}

	var resp matchmakingResponse
	if err := c.conn.ReadJSON(&resp); err != nil {
		return fmt.Errorf("%w: %w", gchessclient.ErrMatchmakingFailed, err)
	}

	if resp.Type != "matched" {
		return fmt.Errorf("%w: %w - got: %s", gchessclient.ErrMatchmakingFailed, ErrExpectMatched, resp.Type)
	}

	// Channel for user to send move in
	moveCh := make(chan string)

	match := domains.Match{
		GameState:   domains.GameState(resp.GameState),
		PlayerState: domains.PlayerState(resp.PlayerState),
		PlayerId:    playerId,
		SessionId:   resp.SessionId,
		MoveCh:      moveCh,
	}
	go c.StartMatch(match)

	return nil
}

func (c *client) StartMatch(match domains.Match) {
	sessionResp := sessionResponse{
		Type:      "session",
		GameState: gameState(match.GameState),
	}
	for {
		app.SyncMatchState(match)
		if sessionResp.GameState.Status != "ACTIVE" {
			// TODO: handle game status
			return
		}
		if match.PlayerState.IsWhiteSide == sessionResp.GameState.IsWhiteTurn {
			if sessionResp.Type == "session" {
				// TODO: valid move
			} else {
				// TODO: invalid move
			}
			move, ok := <-match.MoveCh
			if !ok {
				return
			}

			err := c.conn.WriteJSON(moveRequest{
				Action: "move",
				Data: moveData{
					PlayerId:  match.PlayerId,
					SessionId: match.SessionId,
					Move:      move,
				},
			})
			if err != nil {
				panic(err)
			}
		}
		if err := c.conn.ReadJSON(&sessionResp); err != nil {
			panic(err)
		}
	}
}
