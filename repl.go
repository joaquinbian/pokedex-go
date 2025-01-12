package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
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

				err := c.callback()

				if err != nil {
					fmt.Errorf("Error executing %v command: %w", c.name, err)
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
