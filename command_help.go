package main

import "fmt"

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, value := range getCommand() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}

	fmt.Println()

	return nil
}
