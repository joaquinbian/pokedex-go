package main

import (
	"fmt"
	"os"

	"github.com/joaquinbian/pokedex-go/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next Location Areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous Location Areas",
			callback:    commandMapb,
		},
	}
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex! \nUsage:")

	for k, value := range commands {
		fmt.Printf("%v: %v\n", k, value.description)
	}

	//os.Exit(0)
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {

	res, err := pokeapi.GetLocationAreas(cfg.nextUrl)

	if err != nil {
		return fmt.Errorf("Error mapCommand: %w", err)
	}

	for _, l := range res.Results {
		fmt.Println(l.Name)
	}

	if res.Next != nil {

		cfg.nextUrl = *res.Next
	}

	if res.Previous != nil {

		cfg.prevUrl = *res.Previous
	}

	return nil

}
func commandMapb(cfg *config) error {

	res, err := pokeapi.GetLocationAreas(cfg.prevUrl)

	if err != nil {
		return fmt.Errorf("Error mapCommand: %w", err)
	}

	for _, l := range res.Results {
		fmt.Println(l.Name)
	}

	if res.Next != nil {

		cfg.nextUrl = *res.Next
	}

	if res.Previous != nil {

		cfg.prevUrl = *res.Previous
	}

	return nil

}
