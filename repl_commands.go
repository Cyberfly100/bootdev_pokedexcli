package main

import (
	"fmt"
	"os"

	"github.com/cyberfly100/bootdev_pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	NextArea     string
	PreviousArea string
}

var cfg = &config{
	NextArea:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	PreviousArea: "",
}

var commands = map[string]cliCommand{}

func registerCommand(name, desc string, cb func(*config) error) {
	commands[name] = cliCommand{name: name, description: desc, callback: cb}
}

func initCommands() {
	registerCommand("exit", "Exit the Pokedex", commandExit)
	registerCommand("help", "Show help information", commandHelp)
	registerCommand("map", "Show next 20 areas", commandMap)
	registerCommand("mapb", "Show the previous 20 areas", commandMapb)
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	response, err := pokeapi.GetAreas(cfg.NextArea)
	if err != nil {
		return fmt.Errorf("failed to fetch areas: %w", err)
	}
	cfg.NextArea = response.Next
	if response.Previous != nil {
		cfg.PreviousArea = *response.Previous
	} else {
		cfg.PreviousArea = ""
	}

	fmt.Printf("Areas:\n")
	for _, area := range response.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.PreviousArea == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	response, err := pokeapi.GetAreas(cfg.PreviousArea)
	if err != nil {
		return fmt.Errorf("failed to fetch areas: %w", err)
	}
	cfg.NextArea = response.Next
	if response.Previous != nil {
		cfg.PreviousArea = *response.Previous
	} else {
		cfg.PreviousArea = ""
	}

	fmt.Printf("Areas:\n")
	for _, area := range response.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}
