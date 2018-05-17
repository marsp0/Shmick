package main

func isPerfectSquare(num int) bool {
	// https://leetcode.com/problems/valid-perfect-square/description/
	var (
		low    = 0
		high   = num
		answer = 0
	)
	for low <= high {
		var mid = (high + low) / 2
		answer = mid * mid
		if answer == num {
			return true
		}
		if answer > num {
			high = mid - 1
		}
		if answer < num {
			low = mid + 1
		}
	}
	return false
}
