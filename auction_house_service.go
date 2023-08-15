package worldofwarcraft

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type AuctionHouseService service

func (s *AuctionHouseService) getListAuctionsEndpoint(realmID int) (*url.URL, error) {
	return s.client.apiURL.Parse(fmt.Sprintf("/data/wow/connected-realm/%d/auctions?locale=%s", realmID, s.client.locale))
}

func (s *AuctionHouseService) ListAuctions(ctx context.Context, realmID int, namespace Namespace) ([]*Auction, error) {
	endpoint, err := s.getListAuctionsEndpoint(realmID)
	if err != nil {
		return nil, fmt.Errorf("failed to get ListAuctions endpoint: %w", err)
	}

	request, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	request.Header.Set("Battlenet-Namespace", namespace.String())

	response, err := s.client.Do(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %v", err)
	}

	defer response.Body.Close()

	body := new(struct {
		Auctions []*Auction `json:"auctions"`
	})

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return body.Auctions, nil
}
