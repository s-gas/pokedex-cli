package commands

import (
	"fmt"
	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
)

func commandExplore(config *config.Config, cache *cache.Cache, locationArea string) error {
	if locationArea == "" {
		fmt.Println("No location area specified")
		return nil
	}
	url := config.Url.JoinPath(locationArea)
	fmt.Println(locationArea)
	fmt.Println(url.Path)
	fmt.Println(url.String())
	// make a request to the url with the updated path
	return nil
}
