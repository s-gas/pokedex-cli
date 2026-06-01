package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/s-gas/pokedex-cli/internal/cache"
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

func commandMap(config *config.Config, cache *cache.Cache) error {
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

func commandMapB(config *config.Config, cache *cache.Cache) error {
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
		fmt.Println("not cached")
		var err error
		rawData, err = makeNewRequest(url)
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

func makeNewRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("makeNewRequest: %w", err)
	}
	defer res.Body.Close()
	rawData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("makeNewRequest: %w", err)
	}
	return rawData, nil
}

func printResults(results []Result) {
	for _, result := range results {
		fmt.Println(result.Name)
	}
}
