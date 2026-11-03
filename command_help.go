package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, value := range cfg.commandRegistry {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}

	fmt.Println()

	return nil
}
