package worldofwarcraft

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// roundTripper is a custom RoundTripper for http.Clients
type roundTripper struct {
	accessToken string
	rateLimiter *rate.Limiter
}

func newRoundTripper(token *token) *roundTripper {
	return &roundTripper{
		accessToken: token.AccessToken,
		rateLimiter: rate.NewLimiter(rate.Every(10*time.Millisecond), 1),
	}
}

// RoundTrip sets recurring headers before making a request, also implementing RoundTripper interface
func (t *roundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	if err := t.rateLimiter.Wait(request.Context()); err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+t.accessToken)
	request.Header.Set("Content-Type", "application/json")
	return http.DefaultTransport.RoundTrip(request)
}
