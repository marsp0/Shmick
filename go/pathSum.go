package main

func hasPathSum(root *TreeNode, sum int) bool {
	// https://leetcode.com/problems/path-sum/description/
	if root == nil {
		return false
	}
	return hasPathSumAux(root, sum)
}

func hasPathSumAux(root *TreeNode, sum int) bool {
	if root != nil {
		if sum-root.Val == 0 {
			if root.Left == nil && root.Right == nil {
				return true
			}
		}
		return hasPathSumAux(root.Left, sum-root.Val) || hasPathSumAux(root.Right, sum-root.Val)
	}
	return false
}
