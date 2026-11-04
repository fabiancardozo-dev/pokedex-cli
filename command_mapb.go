package main

import "fmt"

func commandMapB(cfg *config) error {
	if cfg.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := getLocations(cfg.previousURL)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.nextURL = locations.Next
	cfg.previousURL = locations.Previous

	return nil
}
