package main

import (
	"time"

	"github.com/carlosbueloni/pokedexcli/internal/pokeapi"
)

type Config struct {
	Next          *string
	Previuous     *string
	PokeClient    pokeapi.Client
	caughtPokemon map[string]pokeapi.PokemonResp
}

func main() {
	conf := Config{
		caughtPokemon: map[string]pokeapi.PokemonResp{},
		PokeClient:    pokeapi.NewClient(time.Hour),
	}
	startRepl(&conf)
}
