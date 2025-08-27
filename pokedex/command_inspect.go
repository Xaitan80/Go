package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon to inspect")
	}
	pokemon := args[0]
	val, ok := cfg.Pokedex[pokemon]
	if !ok {
		return fmt.Errorf("you have not cought this pokemin: %s", pokemon)
	} else {
		fmt.Printf("Name: %s\n", val.Name)
		fmt.Printf("Height: %d\n", val.Height)
		fmt.Printf("Weight: %d\n", val.Weight)
		fmt.Println("Stats:")
		for _, element := range val.Stats {
			fmt.Printf("  -%s: %d\n", element.Stat.Name, element.BaseStat)
		}
		fmt.Println("Types:")
		for _, element := range val.Types {
			fmt.Printf("  - %s\n", element.Type.Name)
		}
		return nil
	}
}
