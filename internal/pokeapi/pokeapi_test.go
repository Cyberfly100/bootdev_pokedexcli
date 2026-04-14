package pokeapi

import (
	"testing"
	"time"
)

func TestGetAreas(t *testing.T) {
	client := NewClient(5 * time.Second)
	response, err := client.GetAreas(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response.Count == 0 {
		t.Fatalf("Expected count to be greater than 0")
	}
	if len(response.Results) == 0 {
		t.Fatalf("Expected results to be non-empty")
	}
}
