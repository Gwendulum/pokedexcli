package main

import "fmt"


func commandHelp(cfg *config, userInput string) error {
	fmt.Printf(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex

map: Shows the next 20 locations
mapb: Shows the previous 20 locations
explore: Shows the Pokemon in the area
catch: Attempt to catch a Pokemon
inspect: Inspect a Pokemon in your Pokedex
pokedex: shows the Pokemon you've caught so far
`)
	return nil
}
