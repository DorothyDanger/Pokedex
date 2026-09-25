package main

import (
	"fmt"
)

func commandPokedex(con *config, _ []string) error {

	if len(con.pokedex) == 0 {
		fmt.Printf("Your pokedex is empty!\n")
		return nil
	}

	fmt.Printf("Your Pokedex:\n")
	// print the list of all the pokemon in the pokedex by name
	for pokemon := range con.pokedex {
		fmt.Printf("  %s\n", pokemon)
	}
	return nil
}
