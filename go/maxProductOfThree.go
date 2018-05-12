package main

func maximumProduct(nums []int) int {
	// https://leetcode.com/problems/maximum-product-of-three-numbers/description/
	var (
		sortedNums = mergeSort1(nums)
		lastIndex  = len(nums) - 1
		firstTwo   = sortedNums[0] * sortedNums[1]
		lastTwo    = sortedNums[lastIndex-1] * sortedNums[lastIndex-2]
	)
	if firstTwo*sortedNums[lastIndex] > lastTwo*sortedNums[lastIndex] {
		return firstTwo * sortedNums[lastIndex]
	}
	return lastTwo * sortedNums[lastIndex]
}

func mergeSort1(array []int) []int {
	if len(array) == 1 {
		return array
	}
	return merge1(mergeSort1(array[:len(array)/2]), mergeSort1(array[len(array)/2:]))
}

func merge1(first, second []int) []int {
	var (
		i     = 0
		j     = 0
		array = []int{}
	)
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			array = append(array, first[i])
			i++
		} else {
			array = append(array, second[j])
			j++
		}
	}
	if i < len(first) {
		array = append(array, first[i:]...)
	} else {
		array = append(array, second[j:]...)
	}
	return array
}
