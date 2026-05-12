package main

// This file contains functions.
// It may be refactored in the future into multiple files

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
#Function to clean text:
Split text into words
Trim leading and trailing whitespace from each word
Return a slice of cleaned lowercase words
*/
func cleanInput(text string) []string {
	// New slice to hold cleaned words
	cleanedWords := []string{}

	// Split text into words
	text = strings.ToLower(text)
	words := strings.Fields(text)

	// trim word and add to cleanedWords
	for _, word := range words {
		cleanedWords = append(cleanedWords, strings.TrimSpace(word))
	}

	return cleanedWords
}

func startPokedex() {
	scanner := bufio.NewScanner(os.Stdin)

	//infinite for loop to read user input until they exit
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			splitInput := cleanInput(input)
			cmd, exists := getCommands()[splitInput[0]]
			if exists {
				err := cmd.callback()
				if err != nil {
					fmt.Printf("Error executing command: %v\n", err)
				}
			}
		}

	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	cmds := getCommands()

	for _, cmd := range cmds {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
