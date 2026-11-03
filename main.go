package main

func main() {
	cfg := config{
		commandRegistry: getCommand(),
	}
	startRepl(&cfg)
}
