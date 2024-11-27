package main

import (
	"errors"
	"fmt"
)

func commandMap(c *Config, args ...string) error {
	resp, err := c.PokeClient.ListLocationAreas(c.Next)
	if err != nil {
		return err
	}

	c.Next = resp.Next
	c.Previuous = resp.Previous

	fmt.Println("Location areas:")
	for _, location := range resp.Results {
		fmt.Println("-", location.Name)
	}
	return nil
}

func commandMapBack(c *Config, args ...string) error {
	if c.Previuous == nil {
		return errors.New("you're on the first page")
	}
	resp, err := c.PokeClient.ListLocationAreas(c.Previuous)
	if err != nil {
		return err
	}

	c.Next = resp.Next
	c.Previuous = resp.Previous

	fmt.Println("Location areas:")
	for _, location := range resp.Results {
		fmt.Println("-", location.Name)
	}
	return nil
}
