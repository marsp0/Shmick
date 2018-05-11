package main

import (
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	// https://leetcode.com/problems/relative-ranks/description/
	var (
		numsSorted = mergeSort(nums)
		hashTable  = map[int]int{}
		result     = []string{}
	)
	for index, value := range numsSorted {
		hashTable[value] = index
	}
	for _, value := range nums {
		if value == numsSorted[len(numsSorted)-1] {
			result = append(result, "Gold Medal")
		} else if value == numsSorted[len(numsSorted)-2] {
			result = append(result, "Silver Medal")
		} else if value == numsSorted[len(numsSorted)-3] {
			result = append(result, "Bronze Medal")
		} else {
			result = append(result, strconv.Itoa(len(nums)-hashTable[value]))
		}
	}
	return result
}

func mergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}
	return merge(mergeSort(array[:len(array)/2]), mergeSort(array[len(array)/2:]))
}

func merge(first, second []int) []int {
	var (
		i, j     = 0, 0
		newArray = []int{}
	)

	for i < len(first) && j < len(second) {
		if first[i] <= second[j] {
			newArray = append(newArray, first[i])
			i++
		} else {
			newArray = append(newArray, second[j])
			j++
		}
	}
	if i < len(first) {
		newArray = append(newArray, first[i:]...)
	} else {
		newArray = append(newArray, second[j:]...)
	}
	return newArray
}
