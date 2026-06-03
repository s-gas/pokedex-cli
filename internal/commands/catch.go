package commands

import (
	"fmt"
	"errors"
	"encoding/json"
	
	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
	"github.com/s-gas/pokedex-cli/internal/request"
)

func commandCatch(config *config.Config, cache *cache.Cache, argument string) error {
	if argument == "" {
		fmt.Println("No Pokémon specified")
		return nil
	}
	url := config.UrlPokemon.JoinPath(argument)
	rawData, err := request.Do(url.String())
	if errors.Is(err, request.NotFound) {
		fmt.Println("Pokemon not found")
		return nil
	}
	if err != nil {
		return fmt.Errorf("commandCatch: %w", err)
	}
	var pokemon Pokemon 
	if err = json.Unmarshal(rawData, &pokemon); err != nil {
		return fmt.Errorf("commandCatch: %w", err)
	}
	catch(pokemon)
	return nil
}

func catch(pokemon Pokemon) {
	fmt.Printf("Throwing a Pokeball at %s\n", pokemon.Name)
}
