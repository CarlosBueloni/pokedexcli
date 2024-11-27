package main

import "os"

func commandExit(c *Config, arg string) error {
	defer os.Exit(0)
	return nil
}
