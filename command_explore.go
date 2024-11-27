package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a loctaion name")
	}

	resp, err := c.PokeClient.ListEncounters(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", resp.Location.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}
	return nil

}
