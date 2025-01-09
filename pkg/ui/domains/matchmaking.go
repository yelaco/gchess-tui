package domains

type MatchCondition struct {
	// TODO: implement
}

type GameState struct {
	Status      string
	BoardFen    string
	IsWhiteTurn bool
}

type PlayerState struct {
	IsWhiteSide bool
}

type Match struct {
	// TODO: Add user and opponent in match info
	// User
	// Opponent
	GameState   GameState
	PlayerState PlayerState
	PlayerId    string
	SessionId   string
	MoveCh      chan string
	ErrorCh     chan error
}
