package worldofwarcraft

import "net/http"

// Client is a client for World of Warcraft APIs
type Client struct {
	common *service

	AuctionHouse *AuctionHouseService
	Items        *ItemsService
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
			Transport: newRoundTripper(token),
		},
		options: options,
		token:   token,
	}
	client.AuctionHouse = (*AuctionHouseService)(client.common)
	client.Items = (*ItemsService)(client.common)
	client.Realms = (*RealmsService)(client.common)
	return client, nil
}
