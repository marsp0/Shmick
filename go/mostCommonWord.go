package main

import (
	"strings"
)

func mostCommonWord(paragraph string, banned []string) (result string) {
	// https://leetcode.com/problems/most-common-word/description/
	var (
		wordCounter = map[string]int{}
		count       = 0
	)
	var wordList = strings.Split(paragraph, " ")
	for index := range wordList {
		wordList[index] = strings.ToLower(strings.Trim(wordList[index], "!?',;."))
		wordCounter[wordList[index]]++
	}
	for _, value := range banned {
		delete(wordCounter, value)
	}
	for key, value := range wordCounter {
		if value > count {
			result = key
			count = value
		}
	}

	return result
}
