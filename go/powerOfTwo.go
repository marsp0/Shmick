package main

func isPowerOfTwo(n int) bool {
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
		newNumber *= 2
	}
}
