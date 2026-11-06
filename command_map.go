package main

import "fmt"

func commandMap(cfg *config) error {
	respLocations, err := cfg.pokeapiClient.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = respLocations.Next
	cfg.previousURL = respLocations.Previous

	for _, loc := range respLocations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
