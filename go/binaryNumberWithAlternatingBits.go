package main

import (
	"strconv"
)

func hasAlternatingBits(n int) bool {
	// https://leetcode.com/problems/binary-number-with-alternating-bits/description/
	var binaryForm = strconv.FormatInt(int64(n), 2)
	for index, digit := range binaryForm {
		if index > 0 {
			if string(digit) == string(binaryForm[index-1]) {
				return false
			}
		}
	}
	return true
}
