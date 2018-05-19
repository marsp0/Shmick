package main

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	// https://leetcode.com/problems/repeated-substring-pattern/description/
	for i := len(s) / 2; i > 0; i-- {
		var result = strings.Repeat(s[:i], len(s)/len(s[:i]))
		if result == s {
			return true
		}
	}
	return false
}
