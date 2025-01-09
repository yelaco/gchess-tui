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

	// Channel to notify user about change in game state
	errorCh := make(chan error)

	match := domains.Match{
		GameState: domains.GameState{
			Status:      resp.GameState.Status,
			BoardFen:    resp.GameState.BoardFen,
			IsWhiteTurn: resp.GameState.IsWhiteTurn,
		},
		PlayerState: domains.PlayerState{
			IsWhiteSide: resp.PlayerState.IsWhiteSide,
		},
		PlayerId:  playerId,
		SessionId: resp.SessionId,
		MoveCh:    moveCh,
		ErrorCh:   errorCh,
	}
	app.NewMatch(match)
	go c.StartMatch(resp.GameState, match)

	return nil
}

func (c *client) StartMatch(currentState gameState, match domains.Match) {
	sessionResp := sessionResponse{
		Type:      "session",
		GameState: currentState,
	}
	for {
		if sessionResp.GameState.Status != "ACTIVE" {
			match.ErrorCh <- nil
			return
		}
		if match.PlayerState.IsWhiteSide == sessionResp.GameState.IsWhiteTurn {
			if sessionResp.Type != "session" {
				match.ErrorCh <- ErrInvalidMove
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
				match.ErrorCh <- err
			}
		}
		if err := c.conn.ReadJSON(&sessionResp); err != nil {
			match.ErrorCh <- err
		}
		match.GameState = domains.GameState{
			Status:      sessionResp.GameState.Status,
			BoardFen:    sessionResp.GameState.BoardFen,
			IsWhiteTurn: sessionResp.GameState.IsWhiteTurn,
		}
		app.SyncMatch(match)
		match.ErrorCh <- nil
	}
}

func (c *client) Close() error {
	return c.conn.Close()
}
