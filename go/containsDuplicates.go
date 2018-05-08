package main

func containsDuplicate(nums []int) bool {
	// https://leetcode.com/problems/contains-duplicate/description/
	var (
		hashTable = map[int]int{}
	)
	for i := 0; i < len(nums); i++ {
		hashTable[nums[i]]++
		if hashTable[nums[i]] > 1 {
			return true
		}
	}

	return false
}
