package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon")
	}
	pokemon := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	num := r.Intn(251)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if num > pokemonResp.BaseExperience {
		cfg.Pokedex[pokemonResp.Name] = pokemonResp
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
