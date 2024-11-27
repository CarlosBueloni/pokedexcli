package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}
	pokemon, err := c.PokeClient.ListPokemonInfo(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a pokeball at %v...\n", pokemon.Name)
	catch_rate := 100.0 / float64(pokemon.BaseExperience)
	chance := rand.Float64()
	if chance < catch_rate {
		fmt.Printf("%v was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}
