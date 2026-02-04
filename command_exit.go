package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, userInput string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}