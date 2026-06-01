package commands

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io"

	"github.com/s-gas/pokedex-cli/internal/config"
)

type Result struct {
	Name		string		`json:"name"`
	Url			string		`json:"url"`
}

type Locations struct {
	Next		string		`json:"next"`
	Prev		string		`json:"previous"`
	Results	[]Result	`json:"results"`
}

func commandMap(config *config.Config) error {
	if config.Next == "" {
		fmt.Println("No next locations")
		return nil
	}
	res, err := http.Get(config.Next)
	if err != nil {
		return fmt.Errorf("commandMap: %w", err)
	}
	defer res.Body.Close()
	rawData, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("commandMap: %w", err)
	}
	var locations Locations
	if err = json.Unmarshal(rawData, &locations); err != nil {
		return fmt.Errorf("commandMap: %w", err)
	}
	printResults(locations.Results)	
	config.Prev = config.Next
	config.Next = locations.Next
	return nil
}

func commandMapB(config *config.Config) error {
	if config.Prev == "" {
		fmt.Println("No previous locations")
		return nil
	}
	res, err := http.Get(config.Prev)
	if err != nil {
		return fmt.Errorf("commandMap: %w", err)
	}
	defer res.Body.Close()
	rawData, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("commandMap: %w", err)
	}
	var locations Locations
	if err = json.Unmarshal(rawData, &locations); err != nil {
		return fmt.Errorf("commandMap: %w", err)
	}
	printResults(locations.Results)
	config.Next = config.Prev
	config.Prev = locations.Prev
	return nil
}

func printResults(results []Result) {
	for _, result := range results {
		fmt.Println(result.Name)
	}
}
