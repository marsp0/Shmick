package main

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	// https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
	var (
		array  = []int{}
		minVal = math.MaxInt64
	)
	minDiffInBSTAbsoluteAux(root, &array)
	for i := 1; i < len(array); i++ {
		var diff = array[i] - array[i-1]
		if diff < 0 {
			diff = -diff
		}
		if diff < minVal {
			minVal = array[i] - array[i-1]
		}
	}
	return minVal
}

func minDiffInBSTAbsoluteAux(root *TreeNode, array *[]int) {
	if root != nil {
		minDiffInBSTAbsoluteAux(root.Left, array)
		*array = append(*array, root.Val)
		minDiffInBSTAbsoluteAux(root.Right, array)
	}
}
