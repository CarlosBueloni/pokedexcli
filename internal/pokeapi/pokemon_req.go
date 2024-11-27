package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonInfo(pokemonName string) (PokemonResp, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseURL + endpoint

	dat, ok := c.cache.Get(fullUrl)
	if ok {
		pokemonResp := PokemonResp{}
		err := json.Unmarshal(dat, &pokemonResp)
		if err != nil {
			return PokemonResp{}, err
		}

		return pokemonResp, nil

	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return PokemonResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}

	if resp.StatusCode > 399 {
		return PokemonResp{}, fmt.Errorf("bad status error: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResp{}, err
	}

	pokemonResp := PokemonResp{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}

	c.cache.Add(fullUrl, dat)

	return pokemonResp, nil

}
