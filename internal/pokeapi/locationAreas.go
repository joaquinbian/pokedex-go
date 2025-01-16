package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

func GetLocationAreas(pageUrl *string) (LocationAreaResponse, error) {
	//para la primera vez
	baseUrl := BaseUrl + LocationAreas

	if pageUrl != nil {
		baseUrl = *pageUrl
	}

	res, err := http.Get(baseUrl)

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
