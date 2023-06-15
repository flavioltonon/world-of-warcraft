package worldofwarcraft

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
