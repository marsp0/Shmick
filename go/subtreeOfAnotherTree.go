package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	return isSubtreeAux(s, t, false)
}

func isSubtreeAux(first *TreeNode, second *TreeNode, started bool) bool {
	// https://leetcode.com/problems/subtree-of-another-tree/description/
	var result = false
	if first == nil && second == nil {
		if started {
			return true
		}
		return false
	}
	if first == nil || second == nil {
		return false
	}
	if first.Val == second.Val {
		result = isSubtreeAux(first.Left, second.Left, true) && isSubtreeAux(first.Right, second.Right, true)
	}
	if !result {
		if started {
			return false
		}
		result = isSubtreeAux(first.Left, second, false) || isSubtreeAux(first.Right, second, false)
	}
	return result
}
