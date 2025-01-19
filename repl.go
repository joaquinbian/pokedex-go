package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joaquinbian/pokedex-go/internal/pokecache"
)

type config struct {
	prevUrl *string
	nextUrl *string
	cache   pokecache.Cache
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	var cfg = config{
		nextUrl: nil,
		prevUrl: nil,
		cache:   pokecache.NewCache(5 * time.Second),
	}

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := cleanInput(scanner.Text())
			if len(text) > 0 {

				c, ok := commands[text[0]]
				if !ok {
					fmt.Print("Unknown command\n")
					continue
				}

				err := c.callback(&cfg)

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
