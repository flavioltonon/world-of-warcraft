package worldofwarcraft

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"
)

type ItemsService service

func (s *ItemsService) Namespace() Namespace {
	return NewStaticNamespace(s.options.region)
}

func (s *ItemsService) getItemByIDEndpoint(itemID int) string {
	return fmt.Sprintf("%s/data/wow/item/%d?locale=%s", s.options.apiURL, itemID, s.options.locale)
}

func (s *ItemsService) GetItemByID(itemID int) (*Item, error) {
	request, err := http.NewRequest(http.MethodGet, s.getItemByIDEndpoint(itemID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	request.Header.Set("Battlenet-Namespace", s.Namespace().String())

	response, err := s.httpClient.Do(request)
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

func (s *ItemsService) ListItemsByIDs(itemsIDs []int) ([]*Item, error) {
	var (
		g         errgroup.Group
		rwm       sync.RWMutex
		semaphore = make(chan struct{}, 100)
	)

	items := make([]*Item, 0, len(itemsIDs))

	for i := range itemsIDs {
		semaphore <- struct{}{}

		itemID := itemsIDs[i]

		g.Go(func() error {
			defer func() {
				<-semaphore
			}()

			item, err := s.GetItemByID(itemID)
			if err != nil {
				return err
			}

			rwm.Lock()
			items = append(items, item)
			rwm.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return items, nil
}
