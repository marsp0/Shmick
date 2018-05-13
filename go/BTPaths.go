package main

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	// https://leetcode.com/problems/binary-tree-paths/description/
	return binaryTreePathsAux(root, "")
}

func binaryTreePathsAux(root *TreeNode, pathToHere string) (result []string) {
	if root != nil {
		pathToHere += fmt.Sprintf("%d", root.Val)
		if root.Right == nil && root.Left == nil {
			result = append(result, pathToHere)
			return result
		}
		if root.Right == nil && root.Left != nil {
			var path = binaryTreePathsAux(root.Left, pathToHere+"->")
			fmt.Println(path)
			result = append(result, path...)
		} else if root.Right != nil && root.Left == nil {
			var path = binaryTreePathsAux(root.Right, pathToHere+"->")
			result = append(result, path...)
		} else {
			var (
				left  = binaryTreePathsAux(root.Left, pathToHere+"->")
				right = binaryTreePathsAux(root.Right, pathToHere+"->")
			)
			result = append(result, left...)
			result = append(result, right...)
		}
	}
	return result
}
