package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		firsWord := words[0]

		fmt.Printf("Your command was: %s\n", firsWord)

	}
}

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	result := strings.Fields(lowerCaseText) // Split the string around white spaces
	return result
}
