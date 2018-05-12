package main

func intersect(nums1 []int, nums2 []int) (result []int) {
	// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
	var (
		numberTable = map[int]int{}
	)
	for _, value := range nums1 {
		numberTable[value]++
	}
	for _, value := range nums2 {
		if numberTable[value] > 0 {
			result = append(result, value)
			numberTable[value]--
		}
	}
	return result
}
