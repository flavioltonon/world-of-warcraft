package worldofwarcraft

import "net/http"

// Client is a client for World of Warcraft APIs
type Client struct {
	common *service

	AuctionHouse *AuctionHouseService
	Realms       *RealmsService
}

// NewClient creates a new Client
func NewClient(credentials Credentials, optionFuncs ...ClientOptionFunc) (*Client, error) {
	token, err := getToken(credentials)
	if err != nil {
		return nil, err
	}

	options := defaultClientOptions()

	for _, optionFunc := range optionFuncs {
		optionFunc.apply(&options)
	}

	client := new(Client)
	client.common = &service{
		httpClient: &http.Client{
			Transport: &roundTripper{
				accessToken: token.AccessToken,
			},
		},
		options: options,
		token:   token,
	}
	client.AuctionHouse = (*AuctionHouseService)(client.common)
	client.Realms = (*RealmsService)(client.common)
	return client, nil
}

// roundTripper is a custom RoundTripper for http.Clients
type roundTripper struct {
	accessToken string
}

// RoundTrip sets recurring headers before making a request, also implementing RoundTripper interface
func (t *roundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("Authorization", "Bearer "+t.accessToken)
	request.Header.Set("Content-Type", "application/json")
	return http.DefaultTransport.RoundTrip(request)
}
