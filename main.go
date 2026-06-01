package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		words := parseInput(scanner)
		execCommand(words)
	}
}
