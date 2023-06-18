package worldofwarcraft

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const _authenticationURL = "https://oauth.battle.net/token"

// token is an authentication token required to access Blizzard's APIs
type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type" example:"bearer"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

// getToken returns a token in exchange for a given set of credentials.
func getToken(credentials Credentials) (*token, error) {
	if err := credentials.validate(); err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, _authenticationURL, strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(credentials.ClientID, credentials.ClientSecret)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %v", err)
	}

	defer response.Body.Close()

	var body token

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return &body, nil
}
