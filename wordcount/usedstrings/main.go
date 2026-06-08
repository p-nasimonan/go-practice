package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "apple banana apple orange banana apple"
	words := strings.Fields(text)
	wordcount := make(map[string]int)

	for _, word := range words {
		wordcount[word]++
	}

	for word, value := range wordcount {
		fmt.Printf("%s : %d個\n", word, value)
	}
}
