package commands

import (
	"fmt"
	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
)

func commandExplore(config *config.Config, cache *cache.Cache, locationArea string) error {
	fmt.Println(locationArea)
	return nil
}
