package main
import(
	"strings"
	"bufio"
	"os"
	"fmt"
	"pokedexcli/internal/pokecache"
)

type config struct {
	Next			*string	
	Previous		*string
	cache			*pokecache.Cache
	caughtPokemon 	map[string]Pokemon
}

func StartRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := &config{}
	cfg.caughtPokemon = make(map[string]Pokemon, 0)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		currentText := reader.Text()
		cleanText := CleanInput(currentText)
		keyWord := cleanText[0]
		userInput := ""
		if len(cleanText) > 1 {
			userInput = cleanText[1]
		}
		command, ok := getCommands()[keyWord]
		if ok {
			err := command.callback(cfg, userInput)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}

func CleanInput(text string) []string {
	loText := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(loText)
	return words
}

type cliCommand struct {
	name 		string
	description string
	callback	func(*config, string) error
}



func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"map": {
			name:			"map",
			description:	"Displays a map of locations",
			callback:		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Displays last map of locations",
			callback:		commandMapb,
		},
		"explore": {
			name:			"explore",
			description:	"explore an area",
			callback:		commandExplore,
		},
		"catch": {
			name:			"catch",
			description:	"Catch a Pokemon",
			callback:		commandCatch,
		},
		"inspect": {
			name:			"inspect",
			description:	"inspect a Pokemon",
			callback:		commandInspect,
		},
		"pokedex": {
			name:			"pokedex",
			description:	"show a list of caught pokemon",
			callback:		commandPokedex,
		},
	}
}