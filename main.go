package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		words := parseInput(scanner)
		fmt.Printf("Your command was: %s\n", words[0])
	}
}
