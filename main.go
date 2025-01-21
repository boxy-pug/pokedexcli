package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/boxy-pug/pokedexcli/commands"
	"github.com/boxy-pug/pokedexcli/config"
	"github.com/boxy-pug/pokedexcli/utils"
)

func main() {
	config := &config.Config{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		userCommands := utils.CleanInput(input)

		if len(userCommands) == 0 {
			continue
		}

		commandName := userCommands[0]
		config.Args = userCommands[1:]

		command, exists := commands.GetCommands()[commandName]
		if exists {
			if err := command.Callback(config); err != nil {
				fmt.Printf("Error executing command: %v", command.Name)
			}
		} else {
			fmt.Println("Unknown command", commandName)
		}

	}

}
