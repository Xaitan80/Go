package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("No location")
	}
	location := args[0]
	locationResp, err := cfg.pokeapiClient.GetLocation(location)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", location)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Printf("- %v\n", encounter.Pokemon.Name)
	}
	return nil
}
