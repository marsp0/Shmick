package main

func missingNumber(nums []int) int {
	var (
		summation  int
		currentSum int
	)
	for _, value := range nums {
		currentSum += value
	}
	summation = (len(nums) * (len(nums) + 1)) / 2
	return summation - currentSum
}
