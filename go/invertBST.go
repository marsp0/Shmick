package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	// https://leetcode.com/problems/invert-binary-tree/description/
	invertTreeAux(root)
	return root
}

func invertTreeAux(root *TreeNode) {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTreeAux(root.Left)
		invertTreeAux(root.Right)
	}
}
