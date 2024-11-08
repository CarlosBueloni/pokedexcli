package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = locationsResp.Next
	cfg.Previus = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previus == nil {
		return errors.New("You are on the first page")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.Previus)
	if err != nil {
		return err
	}
	cfg.Next = locationsResp.Next
	cfg.Previus = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}
