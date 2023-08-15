package worldofwarcraft

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type ItemsService service

func (s *ItemsService) Namespace() Namespace {
	return NewStaticNamespace(s.client.region)
}

func (s *ItemsService) getGetItemByIDEndpoint(itemID int) (*url.URL, error) {
	return s.client.apiURL.Parse(fmt.Sprintf("/data/wow/item/%d?locale=%s", itemID, s.client.locale))
}

func (s *ItemsService) GetItemByID(ctx context.Context, itemID int) (*Item, error) {
	endpoint, err := s.getGetItemByIDEndpoint(itemID)
	if err != nil {
		return nil, fmt.Errorf("failed to get GetItemByID endpoint: %w", err)
	}

	request, err := s.client.NewRequest(http.MethodGet, endpoint.String(), nil)
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
		return nil, errors.New("item not found")
	}

	var body Item

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return &body, nil
}
