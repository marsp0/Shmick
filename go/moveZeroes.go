package main

func moveZeroes(nums []int) {
	// https://leetcode.com/problems/move-zeroes/
	var (
		zeroIndex   = len(nums)
		numberIndex = 0
	)
	for numberIndex < len(nums) {
		if nums[numberIndex] == 0 {
			if zeroIndex > numberIndex {
				zeroIndex = numberIndex
			}
			numberIndex++
		} else if nums[numberIndex] != 0 {
			if numberIndex > zeroIndex {
				nums[numberIndex], nums[zeroIndex] = nums[zeroIndex], nums[numberIndex]
				numberIndex = zeroIndex
				zeroIndex = len(nums)
			} else {
				numberIndex++
			}
		}
	}
}
