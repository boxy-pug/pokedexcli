package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(s string) []string {
	clean := strings.Fields(strings.ToLower(s))
	return clean
}

func main() {

	config := &Config{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		userCommands := cleanInput(input)

		if len(userCommands) == 0 {
			continue
		}

		command, exists := getCommands()[userCommands[0]]
		if exists {
			if err := command.callback(config); err != nil {
				fmt.Printf("Error executing command: %v", command.name)
			}
		} else {
			fmt.Println("Unknown command", userCommands[0])
		}

	}

}
