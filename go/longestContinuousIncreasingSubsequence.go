package main

func findLengthOfLCIS(nums []int) int {
	// https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/
	if len(nums) > 0 {
		var (
			counter = 1
			index   = 1
			max     = 1
		)
		for index < len(nums) {
			if nums[index] > nums[index-1] {
				counter++
				if counter > max {
					max = counter
				}
			} else {
				counter = 1

			}
			index++
		}
		return max
	}
	return 0
}
