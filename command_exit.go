package main

import "os"

func commandExit(c *Config) error {
	defer os.Exit(0)
	return nil
}
