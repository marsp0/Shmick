package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	// https://leetcode.com/problems/diameter-of-binary-tree/description/
	return diameterOfBinaryTreeAux(root)[0]
}

func diameterOfBinaryTreeAux(root *TreeNode) []int {
	var (
		result = []int{0, 0}
	)
	if root != nil {
		var (
			left  = diameterOfBinaryTreeAux(root.Left)
			right = diameterOfBinaryTreeAux(root.Right)
			temp  = []int{left[1] + right[1], left[0], right[0]}
			max   int
		)
		if left[1] > right[1] {
			result[1] = left[1] + 1
		} else {
			result[1] = right[1] + 1
		}
		for _, val := range temp {
			if val > max {
				max = val
			}
		}
		result[0] = max
	}
	return result
}
