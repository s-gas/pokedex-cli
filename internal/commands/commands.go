package commands

import (
	"fmt"

	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
)

var commands map[string]CliCommand

func init() {
	commands = map[string]CliCommand{
		"exit": { name:     "exit",
			info:     "Exits the Pokedex",
			callback: commandExit,
		},
		"help": {
			name:     "help",
			info:     "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name:     "map",
			info:     "Displays the next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name:     "mapb",
			info:     "Displays the previous 20 locations",
			callback: commandMapB,
		},
		"explore": {
			name:     "explore",
			info:     "Displays the Pokémon that you can encounter in an area",
			callback: commandExplore,
		},
		"catch": {
			name:     "catch",
			info:     "Tries to catch a Pokémon",
			callback: commandCatch,
		},
		"inspect": {
			name:     "inspect",
			info:     "Displays information about a Pokémon from the Pokédex",
			callback: commandInspect,
		},
		"pokedex": {
			name:     "pokedex",
			info:     "Displays the Pokémon in the Pokédex",
			callback: commandPokedex,
		},
	}
}

func Exec(pokedex map[string]Pokemon, words []string, config *config.Config, cache *cache.Cache) error {
	if len(words) == 0 {
		commandHelp(pokedex, config, cache, "")
		return nil
	}
	word := words[0]
	if cmd, ok := commands[word]; ok {
		var argument string
		if len(words) >= 2 {
			argument = words[1]
		}
		if err := cmd.callback(pokedex, config, cache, argument); err != nil {
			return fmt.Errorf("Exec: %w", err)
		}
	} else {
		fmt.Println("Unknown command")
	}
	return nil
}
