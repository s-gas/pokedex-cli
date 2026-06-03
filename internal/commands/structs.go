package commands

import (
	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
)

type CliCommand struct {
	name     string
	info     string
	callback func(pokedex map[string]Pokemon, config *config.Config, cache *cache.Cache, argument string) error
}

// Structures to store json return from "https://pokeapi.co/api/v2/location-area/{name}"
type Location struct {
	Name				string			`json:"name"`
	Encounters	[]Encounter `json:"pokemon_encounters"`	
}

type Encounter struct {
	Pokemon	Pokemon	`json:"pokemon"`
}

type Pokemon struct {
	Name	string	`json:"name"`
}

// Structures to store json returned from "https://pokeapi.co/api/v2/location-area"
type Locations struct {
	Next    string   `json:"next"`
	Prev    string   `json:"previous"`
	Results []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
