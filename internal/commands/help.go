package commands

import (
	"fmt"

	"github.com/s-gas/pokedex-cli/internal/config"
)

func commandHelp(config *config.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.info)
	}
	return nil
}
