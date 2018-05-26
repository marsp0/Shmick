package main

func arrangeCoins(n int) int {
	// https://leetcode.com/problems/arranging-coins/description/
	var (
		counter int
		row     = 1
	)
	if n == 0 {
		return 0
	}
	for n >= row {
		if n-row >= 0 {
			counter++
			n = n - row
			row++

		} else {
			break
		}
	}
	return counter
}
