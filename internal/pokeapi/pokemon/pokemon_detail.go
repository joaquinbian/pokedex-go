package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/joaquinbian/pokedex-go/internal/pokeapi"
)

func GetPokemon(id string) (PokemonDetailsResponse, error) {

	url := pokeapi.BaseUrl + pokeapi.Pokemon + "/" + id

	res, err := http.Get(url)

	if err != nil {
		return PokemonDetailsResponse{}, fmt.Errorf("Error fetching pokemon: %w\n", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	var pokemonDetailResponse PokemonDetailsResponse
	if err := json.Unmarshal(data, &pokemonDetailResponse); err != nil {
		return PokemonDetailsResponse{}, fmt.Errorf("Error parsing pokemon: %w\n", err)

	}

	return pokemonDetailResponse, nil
}
