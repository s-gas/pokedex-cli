package input

import (
	"bufio"
	"fmt"
	"strings"
)

func Parse(scanner *bufio.Scanner) []string {
	printPrompt()
	if scanner.Scan() {
		input := scanner.Text()
		return cleanInput(input)
	}
	return []string{}
}

func printPrompt() {
	fmt.Print("Pokedex > ")
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}
