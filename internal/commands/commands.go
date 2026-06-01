package commands

import (
	"fmt"

	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/config"
)

type CliCommand struct {
	name     string
	info     string
	callback func(config *config.Config, cache *cache.Cache, argument string) error
}

var commands map[string]CliCommand

func init() {
	commands = map[string]CliCommand{
		"exit": {
			name:     "exit",
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
			name:			"explore",
			info:			"Display the Pokémon of an area",
			callback:	commandExplore,
		},
	}
}

func Exec(words []string, config *config.Config, cache *cache.Cache) error {
	if len(words) == 0 {
		commandHelp(config, cache, "")
		return nil
	}
	word := words[0]
	if cmd, ok := commands[word]; ok {
		var argument string
		if len(words) >= 2 {
			argument = words[1]
		}
		if err := cmd.callback(config, cache, argument); err != nil {
			return fmt.Errorf("Exec: %w", err)
		}
	} else {
		fmt.Println("Unknown command")
	}
	return nil
}
