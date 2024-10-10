package player

type Opponent interface {
	GetName() string
	GetRating() int64
	MakeMove()
}

type HumanPlayer struct{}

type BotPlayer struct{}
