package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(c *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		cmd, ok := getCommands()[command]
		if !ok {
			fmt.Println("command not found")
			continue
		}

		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		err := cmd.callback(c, args...)
		if err != nil {
			fmt.Println(err)
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config, arg ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "displays the previous 20 locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "displays all pokemon on the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "tries to catch a pokemon",
			callback:    commandCatch,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
