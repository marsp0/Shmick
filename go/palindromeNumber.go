package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	// https://leetcode.com/submissions/detail/155941415/
	var (
		strNumber = strconv.Itoa(x)
		i         = 0
		j         = len(strNumber) - 1
	)
	for i < j {
		if strNumber[i] == strNumber[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
