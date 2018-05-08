package main

func maxCount(m int, n int, ops [][]int) int {
	// https://leetcode.com/problems/range-addition-ii/description/
	var (
		maxM = m
		maxN = n
	)
	for _, value := range ops {
		if value[0] < maxM {
			maxM = value[0]
		}
		if value[1] < maxN {
			maxN = value[1]
		}
	}
	return maxM * maxN
}
