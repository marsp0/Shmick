package main

func findMaxAverage(nums []int, k int) float64 {
	// https://leetcode.com/problems/maximum-average-subarray-i/description/
	var (
		max            float64
		currentAverage float64
	)
	if k == len(nums) {
		for i := 0; i < len(nums); i++ {
			max += float64(nums[i]) / float64(k)
		}
	} else {
		max = float64(-10000)
		for i := 0; i < len(nums)-k+1; i++ {
			currentAverage = 0.0
			for j := i; j < i+k; j++ {
				currentAverage += float64(nums[j]) / float64(k)
			}
			if currentAverage > max {
				max = currentAverage
			}
		}
	}
	return max
}
