package main

import(
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"time"
	"pokedexcli/internal/pokecache"
)


func commandExplore(cfg *config, userInput string) error {
	if userInput == "" {
		return fmt.Errorf("Please input a location")
	}
	url := "https://pokeapi.co/api/v2/location-area" + "/" + userInput

	if cfg.cache == nil {
		cfg.cache = pokecache.NewCache(5 * time.Second)
	}
	page, exists := cfg.cache.Get(url)	
	if !exists {
		res, err := http.Get(url)
		if err != nil {		
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}	
		if err != nil {
			return err
		}
		cfg.cache.Add(url, body)
		page = body
	}
	
	area := ExploreArea{}
	err := json.Unmarshal(page, &area)
	if err != nil {
    	return err
	}
	fmt.Printf("Found Pokemon:\n")
	for i := 0; i < len(area.PokemonEncounters); i++ {
		fmt.Printf("- %s\n", area.PokemonEncounters[i].Pokemon.Name)
	}

	return nil
}


type ExploreArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}