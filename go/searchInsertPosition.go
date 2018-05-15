package main

func searchInsert(nums []int, target int) int {
	// https://leetcode.com/problems/search-insert-position/description/
	var (
		low  = 0
		high = len(nums) - 1
	)
	for low <= high {
		var mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
