package worldofwarcraft

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuctionHouseService service

func (s *AuctionHouseService) listAuctionsEndpoint(realmID int) string {
	return fmt.Sprintf("%s/data/wow/connected-realm/%d/auctions?locale=%s", s.options.apiURL, realmID, s.options.locale)
}

func (s *AuctionHouseService) ListAuctions(realmID int) ([]*Auction, error) {
	request, err := http.NewRequest(http.MethodGet, s.listAuctionsEndpoint(realmID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	response, err := s.httpClient.Do(request)
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
