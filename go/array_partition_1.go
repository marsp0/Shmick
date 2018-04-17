//https://leetcode.com/problems/array-partition-i/#/description - easy

package main

import "sort"

func arrayPairSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	var result int
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}
	return result
}
