package pokeapi

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// PokeApiResponse represents the response from the PokeAPI
type PokeApiResponse struct {
	Data   PokeData
	Errors []Error
}

// PokeData represents a Pokémon's data
type PokeData struct {
	Name string
	Url  string
}

// Error represents an error from the PokeAPI
type Error struct {
	Message string
}

// FetchPokemonData retrieves data for a Pokémon from the PokeAPI.
func FetchPokemonData(pokemonName string) (PokeApiResponse, bool) {
	url := strings.ReplaceAll(baseURL, "pokemon/"+pokemonName, "")
	fmt.Printf("Fetching data for %s from %s\n", pokemonName, url)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching data:", err)
		return PokeApiResponse{}, false
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		log.Println("API returned an error:", resp.Status)
		return PokeApiResponse{}, false
	}

	data := PokeApiResponse{}
	data.Data = PokeData{Name: pokemonName, Url: url}

	return data, true
}
