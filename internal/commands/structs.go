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

// Structures to store json returned from "https://pokeapi.co/api/v2/location-area/{name}"
type Location struct {
	Name       string      `json:"name"`
	Encounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

// Structures to store json returned from "https://pokeapi.co/api/v2/pokemon/{name}"

type Pokemon struct {
	Name   string  `json:"name"`
	Exp    int     `json:"base_experience"`
	Height int     `json:"height"`
	Weight int     `json:"weight"`
	Stats  []Stats `json:"stats"`
	Types  []Types `json:"types"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type Types struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
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
