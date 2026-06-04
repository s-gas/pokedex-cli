package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
	"github.com/s-gas/pokedex-cli/internal/request"
)

func commandCatch(pokedex map[string]Pokemon, config *config.Config, cache *cache.Cache, argument string) error {
	if argument == "" {
		fmt.Println("No Pokémon specified")
		return nil
	}
	url := config.UrlPokemon.JoinPath(argument)
	rawData, ok := cache.Get(url.String())
	if !ok {
		var err error
		rawData, err = request.Do(url.String())
		if errors.Is(err, request.NotFound) {
			fmt.Println("Pokemon not found")
			return nil
		}
		if err != nil {
			return fmt.Errorf("commandCatch: %w", err)
		}
		cache.Add(url.String(), rawData)
	}
	var pokemon Pokemon
	if err := json.Unmarshal(rawData, &pokemon); err != nil {
		return fmt.Errorf("commandCatch: %w", err)
	}
	catch(pokedex, pokemon)
	return nil
}

func catch(pokedex map[string]Pokemon, pokemon Pokemon) {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if rand.Intn(400) > pokemon.Exp {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
}
