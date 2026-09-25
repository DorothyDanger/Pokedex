package main

/*
This is the command's struct
*/
type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

/*
Map of available commands in the pokedex.
Calling the getCommands function will return this map
Using the key "exit", for example:
cmd, exists := getCommands()["exit"]

	if exists {
		cmd.callback() // This will call the commandExit function
	}
*/
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Enter the name of a location to reveal the Pokemon available in that area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try and catch a Pokemon by name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon in your Pokedex by name",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the list of Pokemon in your Pokedex",
			callback:    commandPokedex,
		},
	}
}
