package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joaquinbian/pokedex-go/internal/pokecache"
	"github.com/joaquinbian/pokedex-go/internal/pokedex"
)

type config struct {
	prevUrl *string
	nextUrl *string
	cache   pokecache.Cache

	pokedex pokedex.Pokedex
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	var cfg = config{
		nextUrl: nil,
		prevUrl: nil,
		cache:   pokecache.NewCache(5 * time.Second),
		pokedex: pokedex.NewPokedex(),
	}

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			stdin := cleanInput(scanner.Text())
			if len(stdin) > 0 {

				c, ok := commands[stdin[0]]
				if !ok {
					fmt.Print("Unknown command\n")
					continue
				}

				err := c.callback(&cfg, stdin[1:]...)

				if err != nil {
					fmt.Print(err)
				}

				continue

			}

		}

	}
}

func cleanInput(text string) []string {
	textTrimed := strings.Trim(text, " ")
	textLowed := strings.ToLower(textTrimed)
	textSplited := strings.Split(textLowed, " ")
	var finalText []string

	for _, w := range textSplited {
		if len(w) > 0 {
			finalText = append(finalText, w)
		}
	}

	return finalText
}
