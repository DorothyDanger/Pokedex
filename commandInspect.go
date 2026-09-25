package main

import (
	"fmt"
)

func commandInspect(con *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a pokemon name to inspect")
	}

	pokemonName := args[0]

	// If pokemon is not in the pokedex, inform the user
	// Otherwise, print the pokemon's details
	if pokemonData, ok := con.pokedex[pokemonName]; ok {
		// Pokemon details
		fmt.Printf("Name: %s\n", pokemonData.Name)
		fmt.Printf("Height: %d\n", pokemonData.Height)
		fmt.Printf("Weight: %d\n", pokemonData.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemonData.Stats {
			fmt.Printf("  %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, t := range pokemonData.Types {
			fmt.Printf("  %s\n", t.Type.Name)
		}
		return nil
	} else {
		return fmt.Errorf("You have not caught %s yet.", pokemonName)
	}
}
