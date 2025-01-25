package main

import (
	"fmt"
	"os"

	locationArea "github.com/joaquinbian/pokedex-go/internal/pokeapi/location_area"
	locationAreaDetail "github.com/joaquinbian/pokedex-go/internal/pokeapi/location_area_detail"
	pokemon "github.com/joaquinbian/pokedex-go/internal/pokeapi/pokemon"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Use exit to exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "this is the command you just ran",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Use map to display the next Location Areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Use mapb to display the previous Location Areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Use explore to see a list of pokemons located at a given location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Use catch to... catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Use inspect to see information about a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Use pokedex to see a list of all the pokemons catched",
			callback:    commandPokedex,
		},
	}
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Print("Welcome to the Pokedex! \nUsage:")

	for k, value := range commands {
		fmt.Printf("%v: %v\n", k, value.description)
	}

	//os.Exit(0)
	return nil
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, args ...string) error {

	res, err := locationArea.GetLocationAreas(cfg.nextUrl, cfg.cache)

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
func commandMapb(cfg *config, args ...string) error {

	if cfg.prevUrl == nil {
		return fmt.Errorf("You are on the first page!\n")
	}

	res, err := locationArea.GetLocationAreas(cfg.prevUrl, cfg.cache)

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

func commandExplore(cfg *config, args ...string) error {
	//fmt.Printf("args en explore: %v\n", args)

	cityName := args[0]
	res, err := locationAreaDetail.GetLocationAreasDetail(cityName, cfg.cache)

	if err != nil {
		return fmt.Errorf("error commandExplore: %w\n", err)
	}

	fmt.Println("Exploring " + cityName + "...")
	fmt.Println("Pokemons found:")
	for _, item := range res.PokemonEncounters {
		fmt.Printf("  - %v\n", item.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {

	pokeName := args[0]

	fmt.Println("Throwing a Pokeball at " + pokeName + "...")

	_, ok := cfg.pokedex.Get(pokeName)

	if ok {
		fmt.Printf("%v was already cached!\n", pokeName)

		return nil
	}

	res, err := pokemon.GetPokemon(pokeName)

	if err != nil {
		return fmt.Errorf("Error command catch: %w", err)
	}

	pokeBaseExp := res.BaseExperience

	if wasPokemonCaught(pokeBaseExp) {
		fmt.Println(pokeName + " was caught!\n")

		cfg.pokedex.Add(pokeName, res)

	} else {
		fmt.Println(pokeName + " escaped...")

	}
	return nil
}

func commandInspect(cfg *config, args ...string) error {

	pokeName := args[0]

	if err := cfg.pokedex.ShowPokemonInfo(pokeName); err != nil {
		return fmt.Errorf("Error inspect command: %w", err)
	}

	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	cfg.pokedex.ShowPokedex()

	return nil
}
