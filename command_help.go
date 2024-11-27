package main

import "fmt"

func commandHelp(c *Config, args ...string) error {
	fmt.Println("Welcome to the pokedex")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range getCommands() {
		fmt.Println(command.name)
		fmt.Println(command.description)
		fmt.Println()
	}
	return nil
}
