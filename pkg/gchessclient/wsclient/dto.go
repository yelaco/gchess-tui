package wsclient

type matchmakingRequest struct {
	Action string          `json:"action"`
	Data   matchmakingData `json:"data"`
}

type matchmakingResponse struct {
	Type        string      `json:"type"`
	SessionId   string      `json:"session_id"`
	GameState   gameState   `json:"game_state"`
	PlayerState playerState `json:"player_state"`
}

type moveRequest struct {
	Action string   `json:"action"`
	Data   moveData `json:"data"`
}

type sessionResponse struct {
	Type      string    `json:"type"`
	GameState gameState `json:"game_state"`
}

type matchmakingData struct {
	PlayerId string `json:"player_id"`
}

type moveData struct {
	PlayerId  string `json:"player_id"`
	SessionId string `json:"session_id"`
	Move      string `json:"move"`
}

type gameState struct {
	Status      string `json:"status"`
	BoardFen    string `json:"board_fen"`
	IsWhiteTurn bool   `json:"is_white_turn"`
}

type playerState struct {
	IsWhiteSide bool `json:"is_white_side"`
}
