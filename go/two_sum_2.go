package main

import (
	"fmt"
)

//Naive approach
func twoSumNaive(numbers []int, target int) []int {
	var toReturn = []int{}
	for i := 0; i < len(numbers); i++ {
		var current = numbers[i]
		var toFind = target - current
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] == toFind {
				toReturn = append(toReturn, i+1)
				toReturn = append(toReturn, j+1)
				break
			} else if numbers[j] > toFind {
				break
			}
		}
		if len(toReturn) > 0 {
			break
		}
	}
	return toReturn
}

func twoSum(numbers []int, target int) []int {
	var toReturn = []int{}
	for i := 0; i < len(numbers); i++ {
		var toFind = target - numbers[i]
		var val, err = binarySearch(numbers, i+1, len(numbers)-1, toFind)
		if err != nil {
			continue
		} else {
			toReturn = append(toReturn, i+1, val+1)
		}
	}
	return toReturn
}

func binarySearch(array []int, low, high, target int) (int, error) {
	for low <= high {
		var mid = (low + high) / 2
		if array[mid] == target {
			return mid, nil
		}
		if array[mid] > target {
			high = mid - 1
		}
		if array[mid] < target {
			low = mid + 1
		}
	}
	return 0, fmt.Errorf(fmt.Sprintf("Target number not in array : %d", target))
}
