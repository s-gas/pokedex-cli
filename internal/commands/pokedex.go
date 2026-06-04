
package commands

import (
	"fmt"

	"github.com/s-gas/pokedex-cli/internal/config"
	"github.com/s-gas/pokedex-cli/internal/cache"
)

func commandPokedex(pokedex map[string]Pokemon, config *config.Config, cache *cache.Cache, argument string) error {
	if len(pokedex) == 0 {
		fmt.Println("Your Pokédex is empty")
		return nil
	}
	fmt.Println("Your Pokédex:")
	for _, v := range pokedex {
		fmt.Printf(" - %v\n", v.Name)
	}
	return nil	
}
