package worldofwarcraft

// ClientOptions are options applyable to a Client
type ClientOptions struct {
	// ref: https://develop.battle.net/documentation/guides/regionality-and-apis
	apiURL string

	// ref: https://develop.battle.net/documentation/world-of-warcraft/guides/localization
	locale string

	// ref: https://develop.battle.net/documentation/guides/regionality-and-apis
	region Region
}

func defaultClientOptions() ClientOptions {
	return ClientOptions{
		apiURL: "https://us.api.blizzard.com",
		locale: "en_US",
		region: US_Region,
	}
}

// ClientOptionFunc is a function capable of modifying ClientOptions
type ClientOptionFunc func(options *ClientOptions)

func (fn ClientOptionFunc) apply(options *ClientOptions) { fn(options) }

// WithAPIURL sets the Client with a custom API URL
func WithAPIURL(apiURL string) ClientOptionFunc {
	return func(options *ClientOptions) { options.apiURL = apiURL }
}

// WithLocale sets the Client with a custom locale
func WithLocale(locale string) ClientOptionFunc {
	return func(options *ClientOptions) { options.locale = locale }
}

// WithRegion sets the Client with a custom region
func WithRegion(region Region) ClientOptionFunc {
	return func(options *ClientOptions) { options.region = region }
}
