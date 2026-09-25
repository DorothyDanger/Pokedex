package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(con *config, args []string) error {
	// Handle no location provided
	if len(args) < 1 {
		return fmt.Errorf("No location name provided.")
	}
	locationName := args[0]

	url := "https://pokeapi.co/api/v2/location-area/" + locationName
	// if a location is provided, check if it exists in the cache
	// print the pokemon available in the location from the cache
	// otherwise make a request to the API
	// store it in the cache and print the pokemon

	// debugging
	// fmt.Printf("Exploring location: %s\n", locationName)
	// fmt.Printf("URL: %s\n", url)

	// handle explored area in cache
	if val, ok := con.cache.Get(url); ok {
		var pokemonEncounters pokemonEncounters
		if err := json.Unmarshal(val, &pokemonEncounters); err != nil {
			return fmt.Errorf("Error unmarshalling cached data: %v", err)
		}
		// Print location areas from the unmarshalled data
		for _, pokemon := range pokemonEncounters.PokemonEncounters {
			fmt.Printf("%s\n", pokemon.Pokemon.Name)
		}
		return nil
	}

	// handle not in cache
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error finding pokemon for location %s: %v", locationName, err)
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d, and\n data: %s", res.StatusCode, data)
	}
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	// Unmarshal the response data
	var pokemonEncounters pokemonEncounters
	if err := json.Unmarshal(data, &pokemonEncounters); err != nil {
		return fmt.Errorf("Error unmarshalling response data: %v", err)
	}
	// Print pokemon found in the area
	fmt.Printf("Pokemon found in %s:\n", locationName)
	for _, pokemon := range pokemonEncounters.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}

	// Add to the cache
	con.cache.Add(url, data)

	return nil
}
