package utils

import (
	"strings"
)

var excludedWords = map[string]bool{
	"a":   true,
	"the": true,
	"of":  true,
	"in":  true,
	"on":  true,
	"and": true,
	"or":  true,
	"to":  true,
	"at":  true,
	"by":  true,
}

func ExtractValidWords(name string) []string {
	words := strings.Fields(name)
	var validWords []string

	for _, word := range words {
		lower := strings.ToLower(word)
		if !excludedWords[lower] {
			validWords = append(validWords, word)
		}
	}

	return validWords
}
