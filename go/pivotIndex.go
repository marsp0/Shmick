package main

func pivotIndex(nums []int) int {
	// https://leetcode.com/problems/find-pivot-index/description/
	var (
		pivot      = -1
		sum        = 0
		runningSum = 0
	)
	for _, value := range nums {
		sum += value
	}
	for i := 0; i < len(nums); i++ {
		if runningSum*2 == sum-nums[i] {
			pivot = i
			break
		}
		runningSum += nums[i]
	}
	return pivot
}
