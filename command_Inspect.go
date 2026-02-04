package main

import(
	"fmt"
)

func commandInspect(cfg *config, userInput string) error {
	fmt.Printf("inspecting %s", userInput)
	pokemon, exists := cfg.caughtPokemon[userInput]
	if !exists {
		return fmt.Errorf("you have not caught that Pokemon yet!")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for i := 0; i < len(pokemon.Stats); i++ {
		fmt.Printf("- %s: %d\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
	}
	/*fmt.Printf("hp: %s", pokemon.Stats[0].BaseStats)
	fmt.Printf("attack: %s", pokemon.Stats[1].BaseStats)
	fmt.Printf("defense: %s", pokemon.Stats[2].BaseStats)
	fmt.Printf("SpAtt: %s", pokemon.Stats[3].BaseStats)
	fmt.Printf("SpDef: %s", pokemon.Stats[4].BaseStats)
	fmt.Printf("SpDef: %s", pokemon.Stats[4].BaseStats)
	*/
	fmt.Printf("Types:\n")
	for i := 0; i < len(pokemon.Types); i++ {
		fmt.Printf("- %s\n", pokemon.Types[i].Type.Name)
	}

	return nil
		
}