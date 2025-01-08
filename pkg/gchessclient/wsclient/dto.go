package wsclient

type matchmakingRequest struct {
	Action string          `json:"action,omitempty"`
	Data   matchmakingData `json:"data,omitempty"`
}

type matchmakingResponse struct {
	Type        string    `json:"type,omitempty"`
	SessionId   string    `json:"session_id,omitempty"`
	GameState   gameState `json:"game_state,omitempty"`
	PlayerState struct {
		IsWhiteSide bool `json:"is_white_side,omitempty"`
	} `json:"player_state,omitempty"`
}

type moveRequest struct {
	Action string   `json:"action,omitempty"`
	Data   moveData `json:"data,omitempty"`
}

type sessionResponse struct {
	Type      string    `json:"type,omitempty"`
	GameState gameState `json:"game_state,omitempty"`
}

type matchmakingData struct {
	PlayerId string `json:"player_id,omitempty"`
}

type moveData struct {
	PlayerId  string `json:"player_id,omitempty"`
	SessionId string `json:"session_id,omitempty"`
	Move      string `json:"move,omitempty"`
}

type gameState struct {
	Status      string `json:"status,omitempty"`
	BoardFen    string `json:"board_fen,omitempty"`
	IsWhiteTurn bool   `json:"is_white_turn,omitempty"`
}
