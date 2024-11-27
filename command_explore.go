package main

import "fmt"

func commandExplore(c *Config, arg string) error {
	resp, err := c.PokeClient.ListEncounters(arg)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon Encounters:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}
	return nil

}
