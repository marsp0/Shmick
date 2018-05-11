package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findTilt(root *TreeNode) int {
	// https://leetcode.com/problems/binary-tree-tilt/description/
	var result = findTiltAux(root)
	return result[1]
}

func findTiltAux(root *TreeNode) (result []int) {
	if root != nil {
		var left = findTiltAux(root.Left)
		var right = findTiltAux(root.Right)
		var newRunningSum = (left[0]) - (right[0])
		if newRunningSum < 0 {
			newRunningSum = -newRunningSum
		}
		return []int{root.Val + left[0] + right[0], newRunningSum + left[1] + right[1]}
	}
	return []int{0, 0}
}
