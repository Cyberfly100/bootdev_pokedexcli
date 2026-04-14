package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	initCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		user_input := cleanInput(scanner.Text())
		command := user_input[0]
		params := user_input[1:]
		if cmd, ok := commands[command]; !ok {
			fmt.Printf("Unknown command: %s\n", command)
		} else {
			if err := cmd.callback(cfg, params); err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	result := strings.Fields(lower)
	return result
}
