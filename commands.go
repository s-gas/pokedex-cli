package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name			string
	info			string
	callback	func() error
}

var commands map[string]cliCommand

func init() { 
	commands = map[string]cliCommand{
		"exit": {
			name:			"exit",
			info:			"Exit the Pokedex",
			callback:	commandExit,
		},
		"help": {
			name:			"help",
			info:			"Displays a help message",
			callback:	commandHelp,
		},
	}
}

func execCommand(words []string) { 
	if len(words) == 0 {
		commandHelp()
		return
	}
	word := words[0]
	if cmd, ok := commands[word]; ok {
		cmd.callback()
	} else {
		fmt.Println("Unknown command")
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.info)
	}
	return nil
}
