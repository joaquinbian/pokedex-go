package pokedex

import (
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
