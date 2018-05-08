package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	// https://leetcode.com/problems/convert-bst-to-greater-tree/description/
	reverseOrder(root, 0)
	return root
}

func reverseOrder(root *TreeNode, parentSum int) int {
	var (
		sum  int
		temp int
	)
	if root != nil {
		sum += reverseOrder(root.Right, parentSum)
		temp = sum
		sum += root.Val
		root.Val += temp
		root.Val += parentSum
		sum += reverseOrder(root.Left, root.Val)
		return sum
	}
	return 0
}
