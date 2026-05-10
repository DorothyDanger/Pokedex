package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//infinite for loop to read user input until they exit
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			splitInput := cleanInput(input)
			fmt.Printf("Your command was: %s\n", splitInput[0])
		}

	}
}
