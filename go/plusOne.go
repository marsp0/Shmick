package main

func plusOne(digits []int) []int {
	// https://leetcode.com/problems/plus-one/description/
	var (
		i          = len(digits) - 1
		toContinue = true
		result     = []int{}
	)
	for toContinue {
		if digits[i] == 9 {
			digits[i] = 0
			i--
			if i == -1 {
				result = append(result, 1)
				result = append(result, digits...)
				digits = result
				break
			}
		} else {
			digits[i]++
			break
		}
	}
	return digits
}
