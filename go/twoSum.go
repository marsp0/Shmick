package main

func twoSum(nums []int, target int) []int {
	// https://leetcode.com/problems/two-sum/description/
	var (
		numberTable = map[int]int{}
		low         = 0
		high        = 0
		toContinue  = true
	)
	for index, value := range nums {
		var secondIndex, inMap = numberTable[value]
		if inMap && target-value == value {
			low = secondIndex
			high = index
			toContinue = false
			break
		}
		numberTable[value] = index
	}
	if toContinue != false {
		for key, value := range numberTable {
			var secondIndex, inMap = numberTable[target-key]
			if inMap && secondIndex != value {
				if secondIndex > value {
					low = value
					high = secondIndex
				} else {
					low = secondIndex
					high = value
				}
				break
			}
		}
	}
	return []int{low, high}
}
