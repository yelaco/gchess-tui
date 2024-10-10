package engineclient

type EngineClient interface {
	IsReady() bool
	SetupEngine() error
	GetBestMove(fen string) (string, error)
	Evaluate(fen string) (int, error)
}
