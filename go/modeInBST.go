package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	// https://leetcode.com/submissions/detail/154889030/
	var (
		valCounter = map[int]int{}
		result     = []int{}
		max        = 0
	)
	findModeAux(root, valCounter)
	for _, count := range valCounter {
		if count > max {
			max = count
		}
	}
	for index, count := range valCounter {
		if count == max {
			result = append(result, index)
		}
	}
	return result
}

func findModeAux(root *TreeNode, valCounter map[int]int) {
	if root != nil {
		findModeAux(root.Left, valCounter)
		valCounter[root.Val]++
		findModeAux(root.Right, valCounter)
	}
	return
}
