package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	// https://leetcode.com/problems/reshape-the-matrix/description/
	if len(nums[0])*len(nums) != r*c {
		return nums
	}
	var (
		rowCounter int
		colCounter int
		result     [][]int
	)
	result = append(result, []int{})
	for _, row := range nums {
		for _, col := range row {
			result[rowCounter] = append(result[rowCounter], col)
			colCounter++
			if colCounter == c {
				if rowCounter < r-1 {
					result = append(result, []int{})
				}
				rowCounter++
				colCounter = 0
			}
		}
	}
	return result
}
