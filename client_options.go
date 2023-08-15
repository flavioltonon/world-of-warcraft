package worldofwarcraft

import (
	"fmt"
	"net/url"
)

// ClientOptions are options applyable to a Client
type ClientOptions struct {
	// ref: https://develop.battle.net/documentation/guides/regionality-and-apis
	apiURL *url.URL

	// ref: https://develop.battle.net/documentation/guides/using-oauth
	oauthTokenURL *url.URL

	credentials *Credentials

	// ref: https://develop.battle.net/documentation/world-of-warcraft/guides/localization
	locale Locale

	// ref: https://develop.battle.net/documentation/guides/regionality-and-apis
	region Region
}

func defaultClientOptions() ClientOptions {
	return ClientOptions{
		apiURL:        mustParseURL("https://us.api.blizzard.com"),
		oauthTokenURL: mustParseURL("https://oauth.battle.net/token"),
		locale:        en_US_Locale,
		region:        US_Region,
	}
}

func mustParseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(fmt.Errorf("failed to parse url \"%s\": %w", s, err))
	}

	return u
}

// ClientOptionFunc is a function capable of modifying ClientOptions
type ClientOptionFunc func(options *ClientOptions)

func (fn ClientOptionFunc) apply(options *ClientOptions) { fn(options) }

// WithAPIURL sets the Client with a custom API URL
func WithAPIURL(apiURL *url.URL) ClientOptionFunc {
	return func(options *ClientOptions) { options.apiURL = apiURL }
}

// WithCredentials sets the Client with credentials
func WithCredentials(credentials *Credentials) ClientOptionFunc {
	return func(options *ClientOptions) { options.credentials = credentials }
}

// WithLocale sets the Client with a custom locale
func WithLocale(locale Locale) ClientOptionFunc {
	return func(options *ClientOptions) { options.locale = locale }
}

// WithOAuthTokenURL sets the Client with a custom OAuth token URL
func WithOAuthTokenURL(oauthTokenURL *url.URL) ClientOptionFunc {
	return func(options *ClientOptions) { options.oauthTokenURL = oauthTokenURL }
}

// WithRegion sets the Client with a custom region
func WithRegion(region Region) ClientOptionFunc {
	return func(options *ClientOptions) { options.region = region }
}
