package main

import (
	"math"
)

func dominantIndex(nums []int) int {
	// https://leetcode.com/problems/largest-number-at-least-twice-of-others/description/
	var (
		result int
		max    = math.MinInt64
	)
	for index, val := range nums {
		if val > max {
			max = val
			result = index
		}
	}
	for index, val := range nums {
		if index != result {
			if 2*val > max {
				result = -1
			}
		}
	}
	return result
}
