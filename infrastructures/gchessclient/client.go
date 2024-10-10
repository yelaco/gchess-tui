package gchessclient

type GchessClient interface {
	MatchMaking() error
}

type client struct{}

func NewGchessClient() GchessClient {
	return &client{}
}

func (c *client) MatchMaking() error {
	return nil
}
