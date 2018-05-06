package main

import (
	"strings"
)

func toGoatLatin(S string) string {
	// https://leetcode.com/problems/goat-latin/description/
	var (
		words  = strings.Split(S, " ")
		result = []string{}
	)
	for index, word := range words {
		var newWord = transformWord(word, index)
		result = append(result, newWord)
	}
	return strings.Join(result, " ")
}

func transformWord(word string, index int) string {
	var vowels = []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	var isVowel bool
	for _, vowel := range vowels {
		if string(word[0]) == vowel {
			isVowel = true
		}
	}
	if !isVowel {
		word = word[1:] + string(word[0])
	}
	word += "ma"
	for i := 1; i < index+2; i++ {
		word += "a"
	}
	return word
}
