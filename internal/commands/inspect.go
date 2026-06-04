package commands

import (
	"fmt"

	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
)

func commandInspect(pokedex map[string]Pokemon, config *config.Config, cache *cache.Cache, argument string) error {
	if argument == "" {
		fmt.Println("No Pokémon specified")
		return nil
	}
	pokemon, ok := pokedex[argument]
	if !ok {
		fmt.Println("Unknown Pokémon!")
		return nil
	}
	printInfo(pokemon)
	return nil
}

func printInfo(pokemon Pokemon) {
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	printStats(pokemon.Stats)
	printTypes(pokemon.Types)
}

func printStats(stats []Stats) {
	fmt.Println("Stats:")
	for _, s := range stats {
		fmt.Printf(" - %v: %v\n", s.Stat.Name, s.BaseStat)
	}
}

func printTypes(types []Types) {
	fmt.Println("Types:")
	for _, t := range types {
		fmt.Printf(" - %v\n", t.Type.Name)
	}
}


