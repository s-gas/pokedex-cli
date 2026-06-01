package commands

import (
	"net/http"
	"fmt"
	"encoding/json"

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
	res, err := http.Get(config.Url)
	if err != nil {
		return fmt.Errorf("commandMap: %w", err)
	}
	defer res.Body.Close()
	var locations Locations
	if err = json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return fmt.Errorf("commandMap: %w", err)
	}
	fmt.Println(locations.Next)
	return nil
}
