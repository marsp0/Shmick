package main

func findErrorNums(nums []int) []int {
	// https://leetcode.com/problems/set-mismatch/description/
	var (
		numberTable = map[int]int{}
		sum         = 0
		max         = len(nums)
		duplicate   = 0
	)
	for _, value := range nums {
		sum += value
		numberTable[value]++
		if numberTable[value] > 1 {
			duplicate = value
		}
	}
	return []int{duplicate, (max*(max+1)/2 - sum + duplicate)}
}
