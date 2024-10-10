package matchmaking

import "github.com/yelaco/gchess-tui/domains/entities/player"

type MatchType string

const (
	PvP        MatchType = "Play"
	PvE        MatchType = "Play Bots"
	Tournament MatchType = "Tournament"
	Variants   MatchType = "Variants"
)

type Opponent interface {
	MakeMove()
}

type PvECondition struct {
	UserRating int64
}

type PvPCondition struct{}

type MatchInfo struct {
	Opponent player.Opponent
}
