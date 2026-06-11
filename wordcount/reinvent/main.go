package main

import (
	"fmt"
	"strings"
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

func WordCount(s string) map[string]int {

	words := splitwords(s)

	wordcount := make(map[string]int)
	for _, word := range words {
		wordcount[word]++
	}

	return wordcount
}

func compareMaps(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func WordCountTest(wc func(s string) map[string]int) {
	text := "I am apple I am Pen youkan is kawaii pupulie is kawaii"
	wordcount := wc(text)

	words := strings.Fields(text)

	trustwc := make(map[string]int)
	for _, word := range words {
		trustwc[word]++
	}
	if compareMaps(trustwc, wordcount) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func main() {
	WordCountTest(WordCount)
}
