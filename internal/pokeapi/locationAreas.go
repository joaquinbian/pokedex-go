package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const LocationAreasUrl string = "https://pokeapi.co/api/v2/location-area"

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationAreas(url string) (LocationAreaResponse, error) {
	res, err := http.Get(url)

	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Error getting location areas: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Error reading data: %w", err)
	}

	var locationAreas LocationAreaResponse

	if err = json.Unmarshal(data, &locationAreas); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Error getting location areas: %w", err)
	}

	return locationAreas, nil

}
