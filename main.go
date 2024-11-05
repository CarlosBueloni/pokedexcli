package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		command := getCommand()
		value, ok := command[scanner.Text()]
		if ok {
			value.callback()
		} else {
			fmt.Println("Invalid command please try again")
		}
	}
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for k, v := range getCommand() {
		fmt.Printf("%v: %v\n", k, v.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	defer os.Exit(3)
	return nil
}
