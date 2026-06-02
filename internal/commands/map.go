package commands

import (
	"encoding/json"
	"fmt"

	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/request"
	"github.com/s-gas/pokedex-cli/internal/config"
)

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Locations struct {
	Next    string   `json:"next"`
	Prev    string   `json:"previous"`
	Results []Result `json:"results"`
}

func commandMap(config *config.Config, cache *cache.Cache, argument string) error {
	if config.Next == "" {
		fmt.Println("No next locations")
		return nil
	}
	locations, err := getLocations(config.Next, cache)
	if err != nil {
		return fmt.Errorf("commandMap: %w\n", err)
	}
	printResults(locations.Results)
	config.Prev = config.Next
	config.Next = locations.Next
	return nil
}

func commandMapB(config *config.Config, cache *cache.Cache, argument string) error {
	if config.Prev == "" {
		fmt.Println("No previous locations")
		return nil
	}
	locations, err := getLocations(config.Prev, cache)
	if err != nil {
		return fmt.Errorf("commandMap: %w\n", err)
	}
	printResults(locations.Results)
	config.Next = config.Prev
	config.Prev = locations.Prev
	return nil
}

func getLocations(url string, cache *cache.Cache) (Locations, error) {
	rawData, ok := cache.Get(url)
	if !ok {
		var err error
		rawData, err = request.Do(url)
		if err != nil {
			return Locations{}, fmt.Errorf("commandMap: %w", err)
		}
		cache.Add(url, rawData)
	}
	var locations Locations
	if err := json.Unmarshal(rawData, &locations); err != nil {
		return Locations{}, fmt.Errorf("commandMap: %w", err)
	}
	return locations, nil
}


func printResults(results []Result) {
	for _, result := range results {
		fmt.Println(result.Name)
	}
}
