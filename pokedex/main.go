package main

import (
	"time"

	"pokedex/internal/pokeapi" // Use YOUR module name
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
