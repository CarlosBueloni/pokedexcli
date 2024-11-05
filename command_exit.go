package main

import "os"

func commandExit() error {
	defer os.Exit(0)
	return nil
}
