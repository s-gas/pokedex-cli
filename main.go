package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/s-gas/pokedex-cli/internal/cache"
	"github.com/s-gas/pokedex-cli/internal/commands"
	"github.com/s-gas/pokedex-cli/internal/config"
	"github.com/s-gas/pokedex-cli/internal/input"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config.New()
	cache := cache.New(1 * time.Minute)
	for {
		words := input.Parse(scanner)
		if err := commands.Exec(words, &config, cache); err != nil {
			fmt.Fprintf(os.Stdout, "%s\n", err)
		}
	}
}
