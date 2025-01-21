package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/boxy-pug/pokedexcli/config"
	"github.com/boxy-pug/pokedexcli/internal/pokeapi"
	"github.com/boxy-pug/pokedexcli/internal/pokecache"
)

var cache = pokecache.NewCache(10 * time.Second)

func commandMap(config *config.Config) error {
	url := config.NextUrl
	if url == "" {
		url = pokeapi.LocationUrl
	}
	return handleLocations(url, config)
}

func commandMapBack(config *config.Config) error {
	url := config.PrevUrl
	if url == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	return handleLocations(url, config)
}

// orchestrate, call the funcs to get, parse and print
func handleLocations(url string, config *config.Config) error {
	data, err := getCachedOrFetch(url)
	if err != nil {
		return fmt.Errorf("Error getting cache or fetching: %v", err)
	}

	locationResponse, err := parseLocations(data)
	if err != nil {
		return fmt.Errorf("Error parsing locations: %v", err)
	}

	printLocations(locationResponse)

	config.NextUrl = locationResponse.Next
	config.PrevUrl = locationResponse.Previous

	return nil
}

// check cache or make an http request if necessary
func getCachedOrFetch(url string) ([]byte, error) {
	data, found := cache.Get(url)
	if found {
		fmt.Println("Serving from cache")
		return data, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching location areas: %v", err)
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	cache.Add(url, data)

	return data, nil
}

// handle JSON unmarshalling, returning a LocationAreaResponse struct
func parseLocations(data []byte) (pokeapi.LocationAreaResponse, error) {
	var locationResponse pokeapi.LocationAreaResponse
	err := json.Unmarshal(data, &locationResponse)
	if err != nil {
		return locationResponse, fmt.Errorf("error parsing json resp: %v", err)
	}
	return locationResponse, nil
}

// print locations
func printLocations(locationResponse pokeapi.LocationAreaResponse) {
	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}
}
