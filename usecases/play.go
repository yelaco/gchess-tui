package usecases

import (
	"github.com/yelaco/gchess-tui/domains/entities/play"
	"github.com/yelaco/gchess-tui/infrastructures/gchessclient/engineclient"
)

type PlayUsecaseInterface interface {
	GetBestMove() (string, error)
}

type playUsecase struct {
	engineClient engineclient.EngineClient
}

func NewPlayUsecase(engineClient engineclient.EngineClient) PlayUsecaseInterface {
	return &playUsecase{
		engineClient: engineClient,
	}
}

func (u *playUsecase) GetBestMove() (string, error) {
	if u.engineClient == nil {
		return "", play.ErrNilClient
	}
	// TODO: get the fen from app state
	move, err := u.engineClient.GetBestMove("4kb1r/p2rqppp/5n2/1B2p1B1/4P3/1Q6/PPP2PPP/2K4R w k - 0 14")
	if err != nil {
		return "", err
	}

	return move, nil
}
