package main

import "os"

func commandExit(cfg *config) error {
	defer os.Exit(0)
	return nil
}
