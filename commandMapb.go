package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CommandMapb will print the previous 20 locations in the Pokedex
func commandMapb(con *config) error {
	if con.previous == nil {
		return fmt.Errorf("No previous locations available")
	}

	// Check if response is cached
	if val, ok := con.cache.Get(*con.previous); ok {
		var locationAreas locationArea
		if err := json.Unmarshal(val, &locationAreas); err != nil {
			return fmt.Errorf("Error unmarshalling response data: %v", err)
		}
		// Print location areas from the unmarshalled data
		for _, location := range locationAreas.Results {
			fmt.Printf("%s", location.Name)
		}

		// set next and previous in the config
		con.next = locationAreas.Next
		con.previous = locationAreas.Previous

		return nil
	}

	res, err := http.Get(*con.previous)
	// error handling
	if err != nil {
		return fmt.Errorf("Error fetching location areas: %v", err)
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d, and\n data: %s", res.StatusCode, data)
	}
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	//Unmarshal the response data
	var locationAreas locationArea
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return fmt.Errorf("Error unmarshalling response data: %v", err)
	}
	// Print location areas from the unmarshalled data
	for _, location := range locationAreas.Results {
		fmt.Printf("%s", location.Name)
	}

	// Add the JSON response to the cache
	con.cache.Add(*con.previous, data)

	// set next and previous in the config
	con.next = locationAreas.Next
	con.previous = locationAreas.Previous

	return nil
}
