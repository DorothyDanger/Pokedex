package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(con *config, args []string) error {
	// Handle no pokemon name provided
	if len(args) < 1 {
		return fmt.Errorf("No pokemon name provided.")
	}
	pokemonName := args[0]

	// Check if the pokemon is in the pokedex
	if _, ok := con.pokedex[pokemonName]; ok {
		fmt.Printf("%s is already in the Pokedex!\n", pokemonName)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	// Check if the pokemon is in the cache
	if val, ok := con.cache.Get(url); ok {
		var pokemonData pokemon
		if err := json.Unmarshal(val, &pokemonData); err != nil {
			return fmt.Errorf("Error unmarshalling cached data: %v", err)
		}
		// Use base experience with math/rand to determine if the pokemon is caught
		// Higher base experience = harder to catch
		catchChance := rand.Intn(100)
		if catchChance < pokemonData.BaseExperience {
			fmt.Printf("%s escaped!\n", pokemonName)
		} else {
			fmt.Printf("%s was caught!\n", pokemonName)
			// Add the pokemon to the pokedex if caught
			con.pokedex[pokemonName] = pokemonData
		}

	}

	// If not in cache, make a request to the API
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error finding pokemon %s: %v", pokemonName, err)
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
	var pokemonData pokemon
	if err := json.Unmarshal(data, &pokemonData); err != nil {
		return fmt.Errorf("Error unmarshalling response data: %v", err)
	}
	// Use base experience with math/rand to determine if the pokemon is caught
	// Higher base experience = harder to catch
	catchChance := rand.Intn(100)
	if catchChance < pokemonData.BaseExperience {
		fmt.Printf("%s escaped!\n", pokemonName)
		// Add to cache if pokemon escapes so a retry can be made without an API call
		con.cache.Add(url, data)
	} else {
		fmt.Printf("%s was caught!\n", pokemonName)
		// Add the pokemon to the pokedex if caught
		con.pokedex[pokemonName] = pokemonData
	}

	return nil
}
