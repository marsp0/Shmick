package main

func containsNearbyDuplicate(nums []int, k int) bool {
	// https://leetcode.com/problems/contains-duplicate-ii/description/
	if len(nums) > 1 {
		for i := 0; i < len(nums)-1; i++ {
			var goTo int
			if len(nums) < i+k+1 {
				goTo = len(nums)
			} else {
				goTo = i + k + 1
			}
			for j := i + 1; j < goTo; j++ {
				if nums[i] == nums[j] {
					return true
				}
			}
		}
	}
	return false
}
