package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/joaquinbian/pokedex-go/internal/pokeapi"
	"github.com/joaquinbian/pokedex-go/internal/pokecache"
)

func GetLocationAreasDetail(id string, cache pokecache.Cache) (LocationAreaDetailResponse, error) {

	url := fmt.Sprintf("%v%v/%v", pokeapi.BaseUrl, pokeapi.LocationAreas, id)
	dataCached, ok := cache.Get(url)

	if ok {
		var locationDetail LocationAreaDetailResponse

		if err := json.Unmarshal(dataCached, &locationDetail); err != nil {
			return LocationAreaDetailResponse{}, fmt.Errorf("Error getting data from cache: %w", err)
		}

		return locationDetail, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return LocationAreaDetailResponse{}, fmt.Errorf("Error making request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	cache.Add(url, data)

	if err != nil {
		return LocationAreaDetailResponse{}, fmt.Errorf("Error raeding body: %w", err)
	}

	var locationDetail LocationAreaDetailResponse

	if err := json.Unmarshal(data, &locationDetail); err != nil {
		return LocationAreaDetailResponse{}, fmt.Errorf("Error unmarshalling res: %w", err)
	}

	return locationDetail, nil

}
