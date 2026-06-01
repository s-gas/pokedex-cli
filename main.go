package main

import (
	"bufio"
	"os"
	"fmt"

	"github.com/s-gas/pokedex-cli/internal/commands"
	"github.com/s-gas/pokedex-cli/internal/config"
	"github.com/s-gas/pokedex-cli/internal/input"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config.Init()
	for {
		words := input.Parse(scanner)
		if err := commands.Exec(words, &config); err != nil {
			fmt.Fprintf(os.Stdout, "%s\n", err)
		}
	}
}
