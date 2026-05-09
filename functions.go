package main

// This file contains functions.
// It may be refactored in the future into multiple files

import (
	"strings"
)

/*
	Function to clean text:

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
