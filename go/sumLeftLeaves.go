package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	// https://leetcode.com/problems/sum-of-left-leaves/description/
	return sumOfLeftLeavesAux(root, false)
}

func sumOfLeftLeavesAux(root *TreeNode, left bool) (result int) {
	if root != nil {
		if root.Left == nil && root.Right == nil && left {
			return root.Val
		}
		return sumOfLeftLeavesAux(root.Right, false) + sumOfLeftLeavesAux(root.Left, true)
	}
	return result
}
