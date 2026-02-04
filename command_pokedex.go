package main

import(
	"fmt"
)

func commandPokedex(cfg *config, userInput string) error {
	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("You have not caught any Pokemon yet.")
	}
	fmt.Printf("caught Pokemon:\n")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}