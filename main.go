package main

import (
	"fmt"
)

func main() {

	fmt.Print("Hello, World!")
	v := CleanInput("    ")
	a := []string{}
	fmt.Print(v, len(v))
	fmt.Print(a, len(a))
}
