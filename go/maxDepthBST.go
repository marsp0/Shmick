package main

func maxDepth(root *TreeNode) int {
	// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
	if root == nil {
		return 0
	}
	return maxDepthAux(root, 1)
}

func maxDepthAux(root *TreeNode, level int) int {
	var (
		left  int
		right int
		max   int
	)
	if root.Left != nil {
		left = maxDepthAux(root.Left, level+1)
	}
	if root.Right != nil {
		right = maxDepthAux(root.Right, level+1)
	}
	if left > right {
		max = left
	} else {
		max = right
	}
	if max > level {
		return max
	}
	return level
}
