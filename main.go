package main

func main() {
	baseUrl := "https://pokeapi.co/api/v2/location-area"

	cfg := config{
		commandRegistry: getCommand(),
		nextURL:         &baseUrl,
		previousURL:     nil,
	}

	startRepl(&cfg)
}
