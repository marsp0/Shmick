package main

func removeElement(nums []int, val int) int {
	// https://leetcode.com/problems/remove-element/description/
	var (
		availableIndex = len(nums) - 1
		newSize        = len(nums)
		index          = 0
	)
	for index <= availableIndex {
		if nums[index] == val {
			if nums[availableIndex] == val {
				availableIndex--
				newSize--
				continue
			}
			nums[index], nums[availableIndex] = nums[availableIndex], nums[index]
			availableIndex--
			newSize--
		}
		index++
	}
	return newSize
}
