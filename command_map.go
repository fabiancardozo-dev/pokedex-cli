package main

import "fmt"

func commandMap(cfg *config) error {
	locations, err := getLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.previousURL = locations.Previous
	cfg.nextURL = locations.Next

	return nil
}
