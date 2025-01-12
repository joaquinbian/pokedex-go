package main

import (
	"strings"
)

func CleanInput(text string) []string {
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
