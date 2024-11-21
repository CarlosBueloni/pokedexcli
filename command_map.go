package main

import (
	"fmt"
	"log"
)

func commandMap(c *Config) error {
	resp, err := c.PokeClient.ListLocationAreas(c.Next)
	if err != nil {
		log.Fatal(err)
	}

	c.Next = resp.Next
	c.Previuous = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(c *Config) error {
	if c.Previuous == nil {
		fmt.Println("You are on the first page")
		return nil
	}
	resp, err := c.PokeClient.ListLocationAreas(c.Previuous)
	if err != nil {
		log.Fatal(err)
	}

	c.Next = resp.Next
	c.Previuous = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
