package main

import "fmt"

func commandPokedex(c *Config, args ...string) error {
	if len(c.caughtPokemon) == 0 {
		fmt.Println("You haven't caught a pokemon")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.caughtPokemon {
		fmt.Printf("  - %v\n", pokemon.Name)
	}
	return nil
}
