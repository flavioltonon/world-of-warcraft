package worldofwarcraft

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type RealmsService service

func (s *RealmsService) Namespace() Namespace {
	return NewDynamicNamespace(s.client.region)
}

func (s *RealmsService) getGetRealmBySlugEndpoint(realmSlug string) (*url.URL, error) {
	return s.client.apiURL.Parse(fmt.Sprintf("/data/wow/realm/%s?locale=%s", realmSlug, s.client.locale))
}

func (s *RealmsService) GetRealmBySlug(ctx context.Context, realmSlug string) (*Realm, error) {
	endpoint, err := s.getGetRealmBySlugEndpoint(realmSlug)
	if err != nil {
		return nil, fmt.Errorf("failed to get GetRealmBySlug endpoint: %w", err)
	}

	request, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	request.Header.Set("Battlenet-Namespace", s.Namespace().String())

	response, err := s.client.Do(ctx, request)
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
