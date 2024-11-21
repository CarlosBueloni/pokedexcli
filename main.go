package main

import "github.com/carlosbueloni/pokedexcli/internal/pokeapi"

type Config struct {
	Next       *string
	Previuous  *string
	PokeClient pokeapi.Client
}

func main() {
	conf := Config{
		PokeClient: pokeapi.NewClient(),
	}
	startRepl(&conf)
}
