package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := cleanInput(scanner.Text())
			if len(text) > 0 {

				fmt.Printf("Your command was: %v\n", text[0])
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
