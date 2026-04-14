package main

import (
	"fmt"
	"os"

	"github.com/cyberfly100/bootdev_pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokeapi.PokemonResponse
	NextArea      *string
	PreviousArea  *string
}

var commands = map[string]cliCommand{}

func registerCommand(name, desc string, cb func(*config, []string) error) {
	commands[name] = cliCommand{name: name, description: desc, callback: cb}
}

func initCommands() {
	registerCommand("exit", "Exit the pokedex", commandExit)
	registerCommand("help", "Show help information", commandHelp)
	registerCommand("map", "Show the next 20 areas", commandMap)
	registerCommand("mapb", "Show the previous 20 areas", commandMapb)
	registerCommand("explore", "Explore a specific area (usage: explore <area-name>)", commandExplore)
	registerCommand("catch", "Catch a pokemon by name (usage: catch <pokemon-name>)", commandCatch)
	registerCommand("inspect", "Look up the stats of a caught pokemon (usage: inspect <pokemon-name>)", commandInspect)
	registerCommand("pokedex", "List all previously caught pokemon", commandPokedex)
}

func commandExit(cfg *config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, params []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config, params []string) error {
	response, err := cfg.pokeapiClient.GetAreas(cfg.NextArea)
	if err != nil {
		return fmt.Errorf("failed to fetch areas: %w", err)
	}
	cfg.NextArea = response.Next
	cfg.PreviousArea = response.Previous

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *config, params []string) error {
	if cfg.PreviousArea == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	response, err := cfg.pokeapiClient.GetAreas(cfg.PreviousArea)
	if err != nil {
		return fmt.Errorf("failed to fetch areas: %w", err)
	}
	cfg.NextArea = response.Next
	cfg.PreviousArea = response.Previous

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandExplore(cfg *config, params []string) error {
	if len(params) == 0 {
		fmt.Println("Please provide an area name. Usage: explore <area-name>")
		return nil
	}
	areaName := params[0]
	fmt.Printf("Exploring %s...\n", areaName)
	response, err := cfg.pokeapiClient.ExploreArea(areaName)
	if err != nil {
		return fmt.Errorf("failed to explore area: %w", err)
	}
	fmt.Println("Pokemon found:")
	for _, pokemon := range response.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, params []string) error {
	if len(params) == 0 {
		fmt.Println("Please provide a pokemon name. Usage: catch <pokemon-name>")
		return nil
	}
	pokemonName := params[0]
	response, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon info: %w", err)
	}
	success := pokeapi.CatchPokemon(response)
	if success {
		cfg.pokedex[pokemonName] = response
	}
	return nil
}

func commandInspect(cfg *config, params []string) error {
	if len(params) == 0 {
		fmt.Println("Please provide a pokemon name. Usage: inspect <pokemon-name>")
		return nil
	}
	pokemonName := params[0]
	pokemon, exists := cfg.pokedex[pokemonName]
	if !exists {
		fmt.Printf("You haven't caught %s yet!\n", pokemonName)
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, params []string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
