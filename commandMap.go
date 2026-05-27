package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CommandMap will print the next 20 locations in the Pokedex
func commandMap(con *config) error {
	if con.next == nil {
		return fmt.Errorf("No more locations available")
	}

	// Check if the response is in the cache
	if val, ok := con.cache.Get(*con.next); ok {
		var locationAreas locationArea
		if err := json.Unmarshal(val, &locationAreas); err != nil {
			return fmt.Errorf("Error unmarshalling cached data: %v", err)
		}
		// Print location areas from the unmarshalled data
		for _, location := range locationAreas.Results {
			fmt.Printf("%s", location.Name)
		}
		//set previous to the current next url
		//set next to the next url in the response body
		con.previous = con.next
		con.next = locationAreas.Next
		return nil
	}

	res, err := http.Get(*con.next)
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
	con.cache.Add(*con.next, data)

	//set previous to the current next url
	//set next to the next url in the response body
	con.previous = con.next
	con.next = locationAreas.Next

	return nil
}
