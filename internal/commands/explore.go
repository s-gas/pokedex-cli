package commands

import (
	"fmt"
	"encoding/json"
	"errors"

	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
	"github.com/s-gas/pokedex-cli/internal/request"
)

type Pokemon struct {
	Name	string	`json:"name"`
}

type Encounter struct {
	Pokemon	Pokemon	`json:"pokemon"`
}

type Location struct {
	Name				string			`json:"name"`
	Encounters	[]Encounter `json:"pokemon_encounters"`	
}

func commandExplore(config *config.Config, cache *cache.Cache, locationArea string) error {
	if locationArea == "" {
		fmt.Println("No location area specified")
		return nil
	}
	url := config.Url.JoinPath(locationArea)
	rawData, err := request.Do(url.String())
	if errors.Is(err, request.NotFound) {
		fmt.Println("Location area not found")
		return nil
	}
	if err != nil {
		return fmt.Errorf("commandExplore: %w", err)
	}
	var location Location
	if err := json.Unmarshal(rawData, &location); err != nil {
		return fmt.Errorf("commandExplore: %w", err)
	}
	printEncounters(location)
	return nil
}

func printEncounters(location Location) {
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range location.Encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
}
