package commands

import (
	"fmt"

	"github.com/boxy-pug/pokedexcli/config"
)

func commandInspect(config *config.Config) error {
	if len(config.Args) < 1 {
		fmt.Errorf("Missing pokemon argument to inspect")
	}

	p, exists := pokedex[config.Args[0]]
	if !exists {
		return fmt.Errorf("pokemon not in pokedex:")
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Printf("Stats:\n")
	for _, st := range p.Stats {
		fmt.Printf("  -%s: %d\n", st.Stat.Name, st.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typ := range p.Types {
		fmt.Printf("  -%s\n", typ.Type.Name)
	}

	/*
	   	Name: pidgey
	   Height: 3
	   Weight: 18
	   Stats:
	     -hp: 40
	     -attack: 45
	     -defense: 40
	     -special-attack: 35
	     -special-defense: 35
	     -speed: 56
	   Types:
	     - normal
	     - flying
	*/

	return nil

}
