package worldofwarcraft

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RealmsService service

func (s *RealmsService) getRealmBySlugEndpoint(realmSlug string) string {
	return fmt.Sprintf("%s/data/wow/realm/%s?locale=%s", s.options.apiURL, realmSlug, s.options.locale)
}

func (s *RealmsService) GetRealmBySlug(realmSlug string) (*Realm, error) {
	request, err := http.NewRequest(http.MethodGet, s.getRealmBySlugEndpoint(realmSlug), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	response, err := s.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, errors.New("realm not found")
	}

	var body Realm

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return &body, nil
}
