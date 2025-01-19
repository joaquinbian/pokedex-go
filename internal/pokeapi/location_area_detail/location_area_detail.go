package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/joaquinbian/pokedex-go/internal/pokeapi"
)

func GetLocationAreasDetail(id string) (LocationAreaDetailResponse, error) {

	res, err := http.Get(fmt.Sprintf("%v%v/%v", pokeapi.BaseUrl, pokeapi.LocationAreas, id))

	if err != nil {
		return LocationAreaDetailResponse{}, fmt.Errorf("Error making request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationAreaDetailResponse{}, fmt.Errorf("Error raeding body: %w", err)
	}

	var locationDetail LocationAreaDetailResponse

	if err := json.Unmarshal(data, &locationDetail); err != nil {
		return LocationAreaDetailResponse{}, fmt.Errorf("Error unmarshalling res: %w", err)
	}

	return locationDetail, nil

}
