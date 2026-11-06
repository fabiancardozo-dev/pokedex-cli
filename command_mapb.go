package main

import (
	"errors"
	"fmt"
)

func commandMapB(cfg *config) error {
	if cfg.previousURL == nil {
		return errors.New("you're on the first page")
	}

	respLocations, err := cfg.pokeapiClient.ListLocations(cfg.previousURL)
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
