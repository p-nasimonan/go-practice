package main

import (
	"fmt"
)

func splitwords(text string) []string {
	var words []string
	var current_word string
	//inword := true

	for i := 0; i < len(text); i++ {
		char := string(text[i])
		if char == " " {
			if current_word != "" {
				words = append(words, current_word)
				current_word = ""
			}
		} else {
			current_word += char
		}
	}
	if current_word != "" {
		words = append(words, current_word)
	}
	return words
}

func main() {
	text := "I am youkan I speak Japanese I have apple youkan is girls " //english only

	words := splitwords(text)

	wordcount := make(map[string]int)
	for _, word := range words {
		wordcount[word]++
	}

	for word, count := range wordcount {
		fmt.Printf("%s : %d 回\n", word, count)
	}
}
