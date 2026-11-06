package main

import (
	"time"

	"github.com/fabiancardozo-dev/pokedex-cli/internals/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		commandRegistry: getCommand(),
		pokeapiClient:   pokeClient,
	}

	startRepl(&cfg)
}
