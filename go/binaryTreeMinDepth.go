package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	// https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
	if root != nil {
		if root.Left == nil && root.Right == nil {
			return 1
		}
		if root.Left != nil && root.Right != nil {
			var (
				left  = minDepth(root.Left)
				right = minDepth(root.Right)
			)
			if left < right {
				return left + 1
			}
			return right + 1
		} else {
			if root.Left != nil {
				return minDepth(root.Left) + 1
			} else {
				return minDepth(root.Right) + 1
			}
		}
	}
	return 0
}
