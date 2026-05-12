package main

/*
This is the command's struct
*/
type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}
