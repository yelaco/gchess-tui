package wsclient

import "errors"

var (
	ErrWsConnectFailed = errors.New("failed to connect to websocket server")
	ErrExpectMatched   = errors.New("expect message type - want: matched")
	ErrInvalidMove     = errors.New("invalid move")
)
