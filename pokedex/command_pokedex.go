package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Printf("Your Pokedex:\n")
	for _, element := range cfg.Pokedex {
		fmt.Printf(" - %s\n", element.Name)
	}

	return nil
}
