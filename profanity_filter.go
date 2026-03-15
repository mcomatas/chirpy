package main

import (
	"strings"
)

func profanityFilter(chirp string) string {
	//body := strings.ToLower(chirp)
	wordList := strings.Split(chirp, " ")
	cleaned := []string{}
	for _, word := range wordList {
		if strings.Contains(strings.ToLower(word), "kerfuffle") ||
			strings.Contains(strings.ToLower(word), "sharbert") ||
			strings.Contains(strings.ToLower(word), "fornax") {
			cleaned = append(cleaned, "****")
		} else {
			cleaned = append(cleaned, word)
		}
	}

	result := strings.Join(cleaned, " ")
	return result
}
