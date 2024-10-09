package apphttp

import (
	"net/http"
	"time"
)

func NewClient(
	maxIdleConnsPerHost int,
	maxTimeout time.Duration,
) *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConnsPerHost = maxIdleConnsPerHost

	var transport http.RoundTripper
	// if idTokenSource != nil {
	// 	transport = &oauth2.Transport{
	// 		Base:   ktracehttp.NewTransport(t),
	// 		Source: idTokenSource,
	// 	}
	// } else {
	// 	transport = ktracehttp.NewTransport(t)
	// }

	return &http.Client{
		Transport: transport,
		Timeout:   maxTimeout,
	}
}
