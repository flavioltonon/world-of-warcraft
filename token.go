package worldofwarcraft

// Token is an authentication token required to access Blizzard's APIs
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type" example:"bearer"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
