package worldofwarcraft

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

// Client is a client for World of Warcraft APIs
type Client struct {
	credentials Credentials
	options     ClientOptions
}

// NewClient creates a new Client
func NewClient(credentials Credentials, optionFuncs ...ClientOptionFunc) (*Client, error) {
	if err := credentials.validate(); err != nil {
		return nil, err
	}

	options := defaultClientOptions()

	for _, optionFunc := range optionFuncs {
		optionFunc.apply(&options)
	}

	return &Client{
		credentials: credentials,
		options:     options,
	}, nil
}

// ClientOptions are options applyable to a Client
type ClientOptions struct {
	// ref: https://develop.battle.net/documentation/guides/regionality-and-apis
	apiURL string

	// ref: https://develop.battle.net/documentation/world-of-warcraft/guides/namespaces
	namespace string
}

func defaultClientOptions() ClientOptions {
	return ClientOptions{
		apiURL:    "https://us.api.blizzard.com",
		namespace: "static-us",
	}
}

// ClientOptionFunc is a function capable of modifying ClientOptions
type ClientOptionFunc func(options *ClientOptions)

func (fn ClientOptionFunc) apply(options *ClientOptions) { fn(options) }

// WithAPIURL sets the Client with a custom API URL
func WithAPIURL(apiURL string) ClientOptionFunc {
	return func(options *ClientOptions) { options.apiURL = apiURL }
}

// WithNamespace sets the Client with a custom namespace
func WithNamespace(namespace string) ClientOptionFunc {
	return func(options *ClientOptions) { options.namespace = namespace }
}

// Credentials are the credentials required to access the Client functionalities. More info at https://develop.battle.net/access/clients
type Credentials struct {
	ClientID string
	Secret   string
}

func (vo Credentials) validate() error {
	return ozzo.ValidateStruct(&vo,
		ozzo.Field(&vo.ClientID, ozzo.Required),
		ozzo.Field(&vo.Secret, ozzo.Required),
	)
}
