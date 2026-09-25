package main

import (
	"fmt"
)

func commandHelp(con *config, _ []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	cmds := getCommands()

	for _, cmd := range cmds {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
