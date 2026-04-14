package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/cyberfly100/bootdev_pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(10 * time.Minute),
	}
}

func (c *Client) fetchData(address *string, target any) error {
	if address == nil {
		return fmt.Errorf("address is nil")
	}
	u, err := url.Parse(*address)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}
	if u.Scheme == "" {
		*address = baseURL + "/" + *address
	}
	// Check cache first
	body, ok := c.cache.Get(*address)
	if ok {
		//log.Printf("Cache hit for URL: %s", *address)
		err := json.Unmarshal(body, target)
		if err != nil {
			return fmt.Errorf("failed to unmarshal cached data: %w", err)
		}
		return nil
	}
	// Not in cache, fetch from API
	resp, err := c.httpClient.Get(*address)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	c.cache.Add(*address, body)
	//log.Printf("Fetched and cached URL: %s", *address)

	err = json.Unmarshal(body, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}
