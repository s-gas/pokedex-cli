package main

import (
	"bufio"
	"os"

	"github.com/s-gas/pokedex-cli/internal/commands"
	"github.com/s-gas/pokedex-cli/internal/config"
	"github.com/s-gas/pokedex-cli/internal/input"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config.Init()
	for {
		words := input.Parse(scanner)
		commands.Exec(words, &config)
	}
}
