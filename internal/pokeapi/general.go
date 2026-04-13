package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

// Client -
type Client struct {
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
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
	resp, err := c.httpClient.Get(*address)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
