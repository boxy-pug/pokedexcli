package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boxy-pug/pokedexcli/pokeapi"
)

func commandMap(config *Config) error {
	url := config.NextUrl
	if url == "" {
		url = pokeapi.LocationUrl
	}
	return fetchLocations(url, config)
}

func commandMapBack(config *Config) error {
	url := config.PrevUrl
	if url == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	return fetchLocations(url, config)
}

func fetchLocations(url string, config *Config) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching location areas: %w", err)
	}
	defer resp.Body.Close()

	var locationResponse pokeapi.LocationAreaResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locationResponse)
	if err != nil {
		return fmt.Errorf("error parsing json resp: %w", err)
	}

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	config.NextUrl = locationResponse.Next
	config.PrevUrl = locationResponse.Previous

	return nil
}
