package main

func isPowerOfThree(n int) bool {
	// https://leetcode.com/problems/power-of-three/description/
	var newNumber = 1
	if n == 1 {
		return true
	}
	for {
		if newNumber == n {
			return true
		} else if newNumber > n {
			return false
		}
		newNumber *= 3
	}
}
