package main

func intersection(nums1 []int, nums2 []int) (result []int) {
	// https://leetcode.com/problems/intersection-of-two-arrays/description/
	var (
		nums1Counter = map[int]int{}
		nums2Counter = map[int]int{}
	)
	for _, value := range nums1 {
		nums1Counter[value] = 0
	}
	for _, value := range nums2 {
		nums2Counter[value] = 0
	}
	for key := range nums1Counter {
		var _, inMap = nums2Counter[key]
		if inMap {
			result = append(result, key)
		}
	}
	return result
}
