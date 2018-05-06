package main

func isToeplitzMatrix(matrix [][]int) bool {
	// https://leetcode.com/problems/toeplitz-matrix/description/
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
