package commands

import (
	"fmt"
	
	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
)

func commandCatch(config *config.Config, cache *cache.Cache, pokemon string) error {
	if pokemon == "" {
		fmt.Println("No Pokémon specified")
		return nil
	}
	fmt.Println("catch!")
	return nil
}
