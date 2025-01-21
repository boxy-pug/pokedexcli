package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/boxy-pug/pokedexcli/config"
	"github.com/boxy-pug/pokedexcli/internal/pokeapi"
)

var pokedex = make(map[string]pokeapi.Pokemon)

func commandCatch(config *config.Config) error {
	if len(config.Args) < 1 {
		return fmt.Errorf("pass in pokemon name to catch")
	}

	url := fmt.Sprintf("%s%s", pokeapi.PokemonUrl, config.Args[0])
	fmt.Printf("Throwing a Pokeball at %s...\n", config.Args[0])

	data, err := getCachedOrFetch(url)
	if err != nil {
		return fmt.Errorf("error getting cache or fetch pokemon: %v", err)
	}

	pokemon, err := parsePokemonDetail(data)
	if err != nil {
		return fmt.Errorf("error parsing pokemon: %v", err)
	}

	if catchPokemon(pokemon) {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v escaped\n", pokemon.Name)
	}
	/*
		for _, name := range pokedex {
			fmt.Println(name.Name)
		}
	*/

	return nil

}

func parsePokemonDetail(data []byte) (pokeapi.Pokemon, error) {
	var pokemon pokeapi.Pokemon
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokemon, fmt.Errorf("Error unmarshaling pokemon: %v", err)
	}
	return pokemon, nil
}

func catchPokemon(pokemon pokeapi.Pokemon) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(100)
	difficulty := 100 - (pokemon.BaseExperience / 10)
	return randomNumber < difficulty
}
