package main

import (
	"math"
)

func titleToNumber(s string) int {
	// https://leetcode.com/problems/excel-sheet-column-number/description/
	var sum int

	for i := len(s) - 1; i > -1; i-- {
		sum += int(math.Pow(26, float64(len(s)-i-1))) * (int(s[i]) - 64)
	}
	return sum
}
