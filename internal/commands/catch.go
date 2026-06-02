package commands

import (
	"fmt"
	
	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
)

func commandCatch(config *config.Config, cache *cache.Cache, argument string) error {
	fmt.Println("catch!")
	return nil
}
