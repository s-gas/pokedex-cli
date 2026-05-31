package main

import (
	"strings"
	"unicode"
)

func cleanInput(text string) []string {
	words := strings.Fields(text)
	words = capitalizeWords(words)
	return words
}

func capitalizeWords(words []string) []string {
	capitalized := []string{}
	for _, word := range words {
		capitalized = append(capitalized, capitalize(word))
	}
	return capitalized
}

func capitalize(word string) string {
	if word == "" {
		return ""
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}
