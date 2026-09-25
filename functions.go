package main

// This file contains functions.
// It may be refactored in the future into multiple files

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/DorothyDanger/Pokedex/internal/pokecache"
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
	// Cache interval and initialization
	const interval = 50 * time.Second // Adjusted to 50 seconds for testing
	cache := pokecache.NewCache(interval)

	//Create config struct with next field set to the first page of location areas
	startURL := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	con := &config{
		next:     &startURL,
		previous: nil,
		cache:    cache,
		pokedex:  make(map[string]pokemon),
	}

	//infinite for loop to read user input until they exit
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			splitInput := cleanInput(input)
			args := splitInput[1:]
			cmd, exists := getCommands()[splitInput[0]]
			if exists {
				err := cmd.callback(con, args)
				if err != nil {
					fmt.Printf("Error executing command: %v\n", err)
				}
			}
		}

	}
}
