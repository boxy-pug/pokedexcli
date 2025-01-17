package commands

import (
	"fmt"
	"os"

	"github.com/boxy-pug/pokedexcli/config"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			Name:        "Exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "Help",
			Description: "Display help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "Map",
			Description: "Display Names of 20 locations in Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "Map back",
			Description: "Get previous page of locations",
			Callback:    commandMapBack,
		},
	}
}

func commandExit(config *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range GetCommands() {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}
	return nil
}
