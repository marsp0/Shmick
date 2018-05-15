package main

func findLHS(nums []int) int {
	// https://leetcode.com/problems/longest-harmonious-subsequence/description/
	var (
		hashTable = map[int]int{}
		max       = 0
	)
	for _, value := range nums {
		hashTable[value]++
	}
	for key, value := range hashTable {

		var (
			lower, lowerErr   = hashTable[key-1]
			higher, higherErr = hashTable[key+1]
		)
		if lowerErr {
			if value+lower > max {
				max = value + lower
			}
		}
		if higherErr {
			if value+higher > max {
				max = value + higher
			}
		}
	}
	return max
}
