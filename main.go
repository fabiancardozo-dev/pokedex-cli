package main

func main() {
	cfg := config{
		commandRegistry: getCommand(),
		currentURL:      "https://pokeapi.co/api/v2/location-area/",
		previousURL:     "",
	}

	startRepl(&cfg)
}
