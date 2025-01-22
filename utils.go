package main

import (
	"math/rand"
)

func getChanceOfSuccess(n int) int {
	return int((1 - float64(n)/float64(200)) * 100)
}

func wasPokemonCaught(exp int) bool {
	p := getChanceOfSuccess(rand.Intn(4))
	return p >= 90
}
