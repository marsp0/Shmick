package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root != nil {
		var balanced, _ = isBalancedAux(root)
		return balanced
	}
	return true
}

func isBalancedAux(root *TreeNode) (bool, int) {
	// https://leetcode.com/problems/balanced-binary-tree/description/
	if root == nil {
		return true, 0
	}
	var leftBalanced, leftHeight = isBalancedAux(root.Left)
	if !leftBalanced {
		return false, 0
	}
	var rightBalanced, rightHeight = isBalancedAux(root.Right)
	if !rightBalanced {
		return false, 0
	}
	var (
		answer    = rightHeight - leftHeight
		maxHeight = 0
	)
	if answer < 0 {
		answer = -answer
		maxHeight = leftHeight
	} else {
		maxHeight = rightHeight
	}
	if answer <= 1 {
		return true, maxHeight + 1
	}
	return false, maxHeight + 1
}
