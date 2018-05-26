package main

func removeDuplicates(nums []int) int {
	// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
	var (
		newArray = []int{}
	)
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}
	newArray = append(newArray, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != newArray[len(newArray)-1] {
			newArray = append(newArray, nums[i])
		}
	}
	for index, value := range newArray {
		nums[index] = value
	}
	return len(newArray)
}
