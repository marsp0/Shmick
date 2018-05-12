package main

func findShortestSubArray(nums []int) int {
	// https://leetcode.com/problems/degree-of-an-array/description/
	var (
		frequencyCounter = map[int]int{}
		indexCounter     = map[int][]int{}
		maxOccurences    = 0
		minDistance      = len(nums)
	)
	for index, value := range nums {
		frequencyCounter[value]++
		indexCounter[value] = append(indexCounter[value], index)
		if maxOccurences < frequencyCounter[value] {
			maxOccurences = frequencyCounter[value]
		}
	}

	for key, value := range indexCounter {
		if frequencyCounter[key] == maxOccurences {
			if value[len(value)-1]-value[0] < minDistance {
				minDistance = value[len(value)-1] - value[0]
			}
		}
	}
	return minDistance + 1
}
