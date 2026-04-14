package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/pokemon?offset=0&limit=3",
			val: []byte(`{
				"count": 1350,
				"next": "https://pokeapi.co/api/v2/pokemon?offset=3&limit=3",
				"previous": null,
				"results": [
					{
					"name": "bulbasaur",
					"url": "https://pokeapi.co/api/v2/pokemon/1/"
					},
					{
					"name": "ivysaur",
					"url": "https://pokeapi.co/api/v2/pokemon/2/"
					},
					{
					"name": "venusaur",
					"url": "https://pokeapi.co/api/v2/pokemon/3/"
					}
				]
			}`),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area?offset=0&limit=2",
			val: []byte(`{
				"count": 1203,
				"next": "https://pokeapi.co/api/v2/location-area?offset=2&limit=2",
				"previous": null,
				"results": [
					{
					"name": "canalave-city-area",
					"url": "https://pokeapi.co/api/v2/location-area/1/"
					},
					{
					"name": "eterna-city-area",
					"url": "https://pokeapi.co/api/v2/location-area/2/"
					}
				]
			}`),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	url := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	cache.Add(url, []byte("testdata"))

	_, ok := cache.Get(url)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(url)
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
