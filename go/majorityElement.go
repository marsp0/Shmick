package main

func majorityElement(nums []int) (result int) {
	// https://leetcode.com/problems/majority-element/description/
	var (
		letterCounter = map[int]int{}
		resultValue   = 0
	)
	for _, value := range nums {
		letterCounter[value]++
	}
	for key, value := range letterCounter {
		if value > resultValue {
			result = key
			resultValue = value
		}
	}
	return result
}
