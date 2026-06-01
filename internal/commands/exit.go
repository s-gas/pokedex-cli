package commands

import (
	"fmt"
	"os"

	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
)

func commandExit(config *config.Config, cache *cache.Cache, argument string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
