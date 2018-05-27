package main

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	// https://leetcode.com/problems/word-pattern/description/
	var (
		patternToString = map[string]string{}
		stringToPattern = map[string]string{}
		wordList        = strings.Split(str, " ")
	)
	if len(wordList) != len(pattern) {
		return false
	}
	for index, word := range wordList {
		var patternValue, ok = stringToPattern[word]
		if ok {
			if patternValue != string(pattern[index]) {
				return false
			}

		} else {
			var _, ok = patternToString[string(pattern[index])]
			if ok {
				return false
			}
			stringToPattern[word] = string(pattern[index])
			patternToString[string(pattern[index])] = word
		}
	}
	return true
}
