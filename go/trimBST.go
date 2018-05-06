package main

// TreeNode - represents a node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	// https://leetcode.com/problems/trim-a-binary-search-tree/description/
	if root == nil {
		return nil
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	if root.Val < L {
		return root.Right
	}
	if root.Val > R {
		return root.Left
	}
	return root
}
