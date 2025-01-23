package pokedex

import (
	"fmt"
	"sync"

	pokemon "github.com/joaquinbian/pokedex-go/internal/pokeapi/pokemon"
)

type Pokedex struct {
	pokemons map[string]pokemon.PokemonDetailsResponse
	mu       *sync.RWMutex
}

func NewPokedex() Pokedex {
	return Pokedex{
		pokemons: make(map[string]pokemon.PokemonDetailsResponse),
		mu:       &sync.RWMutex{},
	}
}

func (p *Pokedex) Add(name string, pokemon pokemon.PokemonDetailsResponse) {
	p.mu.RLock()

	if _, ok := p.pokemons[name]; !ok {
		p.pokemons[name] = pokemon
	}

	defer p.mu.RUnlock()
}

func (p Pokedex) Get(name string) (pokemon.PokemonDetailsResponse, bool) {
	p.mu.RLock()

	poke, ok := p.pokemons[name]
	defer p.mu.RUnlock()

	if !ok {
		return pokemon.PokemonDetailsResponse{}, false
	}

	return poke, true

}

func (p Pokedex) ShowPokemonInfo(name string) error {
	p.mu.RLock()

	defer p.mu.RUnlock()

	pokemon, ok := p.pokemons[name]

	if !ok {
		return fmt.Errorf("Pokemon %v was not found! Try another one\n", name)
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemon.Types {
		fmt.Printf(" -%v\n", v.Type.Name)
	}

	fmt.Println("")

	return nil
}
