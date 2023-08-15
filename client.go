package worldofwarcraft

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/time/rate"
)

// Client is a client for World of Warcraft APIs
type Client struct {
	apiURL        *url.URL
	credentials   *Credentials
	httpClient    *http.Client
	oauthTokenURL *url.URL
	rateLimiter   *rate.Limiter

	locale Locale
	region Region

	common service

	Authentication *AuthenticationService
	AuctionHouse   *AuctionHouseService
	Items          *ItemsService
	Realms         *RealmsService
}

const _BlizzardAPIMaxRequestsPerSecond = 100

// NewClient creates a new Client
func NewClient(httpClient *http.Client, optionFuncs ...ClientOptionFunc) (*Client, error) {
	options := defaultClientOptions()

	for _, optionFunc := range optionFuncs {
		optionFunc.apply(&options)
	}

	client := &Client{
		apiURL:        options.apiURL,
		credentials:   options.credentials,
		httpClient:    httpClient,
		locale:        options.locale,
		oauthTokenURL: options.oauthTokenURL,
		rateLimiter:   rate.NewLimiter(rate.Limit(_BlizzardAPIMaxRequestsPerSecond), 1),
		region:        options.region,
	}

	client.common.client = client
	client.AuctionHouse = (*AuctionHouseService)(&client.common)
	client.Authentication = (*AuthenticationService)(&client.common)
	client.Items = (*ItemsService)(&client.common)
	client.Realms = (*RealmsService)(&client.common)

	if err := client.Authentication.ValidateCredentials(client.credentials); err != nil {
		return nil, fmt.Errorf("failed to validate credentials: %w", err)
	}

	return client, nil
}

func (c *Client) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	token, err := c.Authentication.GetToken(c.credentials)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+token.AccessToken)
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func (c *Client) Do(ctx context.Context, request *http.Request) (*http.Response, error) {
	if err := c.rateLimiter.Wait(ctx); err != nil {
		return nil, err
	}

	return c.httpClient.Do(request)
}
