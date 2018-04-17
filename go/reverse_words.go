package main

import "strings"

func reverseWords(s string) string {
	var wordList = strings.Split(s, " ")
	for i := 0; i < len(wordList); i++ {
		var temp string
		for j := len(wordList[i]) - 1; j > -1; j-- {
			temp += string(wordList[i][j])
		}
		wordList[i] = temp
	}
	return strings.Join(wordList, " ")
}
