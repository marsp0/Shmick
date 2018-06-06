package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	// https://leetcode.com/problems/length-of-last-word/description/
	var (
		wordList = strings.Split(s, " ")
		index    = len(wordList) - 1
	)
	if len(wordList) > 0 {
		for index >= 0 {
			if wordList[index] != " " && wordList[index] != "" {
				return len(wordList[index])
			}
			index--
		}
	}
	return 0
}
