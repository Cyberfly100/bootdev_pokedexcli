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
	NextArea      *string
	PreviousArea  *string
}

var commands = map[string]cliCommand{}

func registerCommand(name, desc string, cb func(*config, []string) error) {
	commands[name] = cliCommand{name: name, description: desc, callback: cb}
}

func initCommands() {
	registerCommand("exit", "Exit the Pokedex", commandExit)
	registerCommand("help", "Show help information", commandHelp)
	registerCommand("map", "Show next 20 areas", commandMap)
	registerCommand("mapb", "Show the previous 20 areas", commandMapb)
	registerCommand("explore", "Explore a specific area (usage: explore <area-name>)", commandExplore)
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
