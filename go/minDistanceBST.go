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
func minDiffInBST(root *TreeNode) int {
	// https://leetcode.com/problems/minimum-distance-between-bst-nodes/description/
	var (
		array  = []int{}
		minVal = math.MaxInt64
	)
	minDiffInBSTAux(root, &array)
	for i := 1; i < len(array); i++ {
		if array[i]-array[i-1] < minVal {
			minVal = array[i] - array[i-1]
		}
	}
	return minVal
}

func minDiffInBSTAux(root *TreeNode, array *[]int) {
	if root != nil {
		minDiffInBSTAux(root.Left, array)
		*array = append(*array, root.Val)
		minDiffInBSTAux(root.Right, array)
	}
}
