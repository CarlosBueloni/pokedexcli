package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, v := range getCommands() {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	fmt.Println()
	return nil
}
