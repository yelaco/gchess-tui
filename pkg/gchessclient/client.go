package gchessclient

type Client interface {
	Matchmaking() error
	Close() error
}
