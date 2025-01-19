package main

import (
	"fmt"
	"os"

	"github.com/joaquinbian/pokedex-go/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args []string) error
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
		"explore": {
			name:        "explore",
			description: "Shows a list of pokemons located at a given location area",
			callback:    commandExplore,
		},
	}
}

func commandHelp(cfg *config, args []string) error {
	fmt.Print("Welcome to the Pokedex! \nUsage:")

	for k, value := range commands {
		fmt.Printf("%v: %v\n", k, value.description)
	}

	//os.Exit(0)
	return nil
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, args []string) error {

	res, err := pokeapi.GetLocationAreas(cfg.nextUrl, cfg.cache)

	if err != nil {
		return fmt.Errorf("Error mapCommand: %w", err)
	}

	for _, l := range res.Results {
		fmt.Println(l.Name)
	}

	cfg.nextUrl = res.Next

	cfg.prevUrl = res.Previous

	return nil

}
func commandMapb(cfg *config, args []string) error {

	if cfg.prevUrl == nil {
		return fmt.Errorf("You are on the first page!\n")
	}

	res, err := pokeapi.GetLocationAreas(cfg.prevUrl, cfg.cache)

	if err != nil {
		return fmt.Errorf("Error mapCommand: %w", err)
	}

	for _, l := range res.Results {
		fmt.Println(l.Name)
	}

	cfg.nextUrl = res.Next

	cfg.prevUrl = res.Previous

	return nil

}

func commandExplore(cfg *config, args []string) error {
	//fmt.Printf("args en explore: %v\n", args)
	return nil
}
