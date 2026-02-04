package main

import(
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"time"
	"pokedexcli/internal/pokecache"
)

const interval = 5 * time.Second


func commandMap(cfg *config, userInput string) error {
	var url string
	if cfg.Next == nil && cfg.Previous == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else if cfg.Next != nil {
		url = *cfg.Next
	} else {
		return fmt.Errorf("You're on the last page")
	}
	err := generateMapPage(url, cfg)
	return err
}


func commandMapb(cfg *config, userInput string) error {
	var url string
	if cfg.Next == nil && cfg.Previous == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else if cfg.Previous != nil {
		url = *cfg.Previous
	} else {
		return fmt.Errorf("You're on the first page")
	}
	err := generateMapPage(url, cfg)
	return err
}


func generateMapPage(url string, cfg *config) error {
	if cfg.cache == nil {
		cfg.cache = pokecache.NewCache(interval)
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
	location := Location{}
	err := json.Unmarshal(page, &location)
	if err != nil {
    	return err
	}
	for i := 0; i < len(location.Results); i++ {
		fmt.Printf("%s\n", location.Results[i].Name)
	}

	cfg.Next = location.Next
	cfg.Previous = location.Previous

	return nil
}


type Location struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}