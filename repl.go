package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	red    = "\033[31m"
	blue   = "\033[94m"
	yellow = "\033[93m"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commandRegistry map[string]cliCommand
	nextURL         *string
	previousURL     *string
}

type Location struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations, Each subsequent call display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%s%sPokedex > %s", bold, blue, reset)

		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		command, ok := getCommand()[words[0]]

		if !ok {
			fmt.Printf("%s%sUnknown command, type 'help' to see all the commands\n", bold, yellow)
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	result := strings.Fields(lowerCaseText) // Split the string around white spaces
	return result
}

func getLocations(url *string) (Location, error) {
	res, err := http.Get(*url)
	if err != nil {
		return Location{}, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Location{}, errors.New("Request failed with status code: " + string(res.StatusCode))
	}
	if err != nil {
		return Location{}, err
	}

	var locs Location
	err = json.Unmarshal(body, &locs)
	if err != nil {
		return Location{}, err
	}

	return locs, nil
}
