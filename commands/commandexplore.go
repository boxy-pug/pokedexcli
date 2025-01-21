package commands

import (
	"encoding/json"
	"fmt"

	"github.com/boxy-pug/pokedexcli/config"
	"github.com/boxy-pug/pokedexcli/internal/pokeapi"
)

func commandExplore(config *config.Config) error {
	if len(config.Args) < 1 {
		return fmt.Errorf("Please provide a location to explore")
	}

	url := fmt.Sprintf("%s%s", pokeapi.LocationUrl, config.Args[0])

	data, err := getCachedOrFetch(url)
	if err != nil {
		return fmt.Errorf("Error fetching cache or url", err)
	}

	locationDetail, err := parseLocationDetail(data)
	if err != nil {
		return fmt.Errorf("Error parsing location detail", err)
	}

	printPokemon(locationDetail)

	return nil
}

func parseLocationDetail(data []byte) (pokeapi.LocationDetailResponse, error) {
	var locationDetail pokeapi.LocationDetailResponse
	err := json.Unmarshal(data, &locationDetail)
	if err != nil {
		return locationDetail, fmt.Errorf("Error parsing location detail", err)
	}
	return locationDetail, nil

}

func printPokemon(locationDetail pokeapi.LocationDetailResponse) {
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationDetail.PokemonEncounters {
		fmt.Println("- ", encounter.Pokemon.Name)
	}

}
