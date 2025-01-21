package commands

import (
	"fmt"

	"github.com/boxy-pug/pokedexcli/config"
)

func commandPokedex(config *config.Config) error {
	if len(pokedex) < 1 {
		fmt.Println("No pokemon in your pokedex")
		return nil
	}
	fmt.Println("Your Pokemon:")
	for _, pokemon := range pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}
