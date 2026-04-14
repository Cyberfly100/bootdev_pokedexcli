package main

import (
	"time"

	"github.com/cyberfly100/bootdev_pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)

	cfg := &config{
		pokeapiClient: client,
		pokedex:       make(map[string]pokeapi.PokemonResponse),
	}
	startRepl(cfg)
}
