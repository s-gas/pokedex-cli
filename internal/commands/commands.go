package commands

import (
	"fmt"

	"github.com/s-gas/pokedex-cli/internal/config"
)

type CliCommand struct {
	name     string
	info     string
	callback func(config *config.Config) error
}

var commands map[string]CliCommand

func init() {
	commands = map[string]CliCommand{
		"exit": {
			name:     "exit",
			info:     "Exit the Pokedex",
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
	}
}

func Exec(words []string, config *config.Config) {
	if len(words) == 0 {
		commandHelp(config)
		return
	}
	word := words[0]
	if cmd, ok := commands[word]; ok {
		cmd.callback(config)
	} else {
		fmt.Println("Unknown command")
	}
}
