package main

func findMaxConsecutiveOnes(nums []int) int {
	var max int
	var counter int
	for _, v := range nums {
		if v == 1 {
			counter++
		} else if v == 0 {
			if counter > max {
				max = counter
			}
			counter = 0
		}
	}
	if counter > max {
		max = counter
	}
	return max
}
