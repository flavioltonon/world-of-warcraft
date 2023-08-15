package worldofwarcraft

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type AuthenticationService service

// GetToken returns a token in exchange for a given set of credentials.
func (s *AuthenticationService) GetToken(credentials *Credentials) (*Token, error) {
	if err := credentials.validate(); err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, s.client.oauthTokenURL.String(), strings.NewReader("grant_type=client_credentials"))
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

	var body Token

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return &body, nil
}

func (s *AuthenticationService) ValidateCredentials(credentials *Credentials) error {
	if _, err := s.GetToken(credentials); err != nil {
		return ErrInvalidCredentials
	}

	return nil
}
