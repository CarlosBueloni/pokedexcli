package main

import (
	"time"

	"github.com/carlosbueloni/pokedexcli/internal/pokeapi"
)

type Config struct {
	Next       *string
	Previuous  *string
	PokeClient pokeapi.Client
}

func main() {
	conf := Config{
		PokeClient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&conf)
}
