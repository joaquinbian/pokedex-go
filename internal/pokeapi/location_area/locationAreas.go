package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/joaquinbian/pokedex-go/internal/pokeapi"
	"github.com/joaquinbian/pokedex-go/internal/pokecache"
)

func GetLocationAreas(pageUrl *string, cache pokecache.Cache) (LocationAreaResponse, error) {
	//para la primera vez
	baseUrl := pokeapi.BaseUrl + pokeapi.LocationAreas

	if pageUrl != nil {
		baseUrl = *pageUrl
	}

	var locationAreas LocationAreaResponse

	dataCached, ok := cache.Get(baseUrl)

	if ok {
		if err := json.Unmarshal(dataCached, &locationAreas); err != nil {
			return LocationAreaResponse{}, fmt.Errorf("Error getting location areas from cache: %w", err)

		}

		return locationAreas, nil
	}

	res, err := http.Get(baseUrl)

	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Error getting location areas: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	cache.Add(baseUrl, data)

	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Error reading data: %w", err)
	}

	if err = json.Unmarshal(data, &locationAreas); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Error getting location areas: %w", err)
	}

	return locationAreas, nil

}
