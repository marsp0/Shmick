package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	// https://leetcode.com/submissions/detail/154151534/
	if root != nil {
		return isSymmetricAux(root.Left, root.Right)
	}
	return true
}

func isSymmetricAux(left, right *TreeNode) bool {
	if left != nil && right != nil {
		if left.Val == right.Val {
			return isSymmetricAux(left.Left, right.Right) && isSymmetricAux(left.Right, right.Left)
		}
	} else if left == nil && right == nil {
		return true
	}
	return false
}
