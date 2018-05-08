package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) (result string) {
	// https: //leetcode.com/problems/construct-string-from-binary-tree/description/
	var (
		array = []string{}
	)
	tree2strAux(root, &array)
	return strings.Join(array, "")
}

func tree2strAux(root *TreeNode, array *[]string) {
	if root != nil {
		(*array) = append(*array, strconv.Itoa(root.Val))

		if root.Left != nil || root.Right != nil {
			(*array) = append(*array, "(")
			tree2strAux(root.Left, array)
			(*array) = append(*array, ")")
		}
		if root.Right != nil {
			(*array) = append(*array, "(")
			tree2strAux(root.Right, array)
			(*array) = append(*array, ")")
		}

	}
}
