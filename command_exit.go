package main

import "os"

func commandExit(c *Config, args ...string) error {
	defer os.Exit(0)
	return nil
}
