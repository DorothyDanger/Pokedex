package main

// This is the struct for the Pokemon Encounters API response
// unmarshals the JSON response into a Go struct
// for the explore command to display the pokemon available
type pokemonEncounters struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
