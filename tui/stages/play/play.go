package play

type GameMoveMsg struct {
	Fen  string
	Move string
}

type GameUpdateMsg struct {
	Accepted    bool
	Fen         string
	IsWhiteTurn bool
}
