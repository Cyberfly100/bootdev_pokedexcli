package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

func fetchData(url string, target any) error {
	resp, err := http.Get(url)
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
