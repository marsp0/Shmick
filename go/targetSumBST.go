package main

func findTarget(root *TreeNode, k int) bool {
	// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
	var (
		numberTable = map[int]int{}
	)
	findTargetAux(root, numberTable, k)
	for key, value := range numberTable {
		if key != value {
			var _, inMap = numberTable[value]
			if inMap {
				return true
			}
		}
	}
	return false
}

func findTargetAux(root *TreeNode, numberTable map[int]int, target int) {
	if root != nil {
		findTargetAux(root.Left, numberTable, target)
		numberTable[root.Val] = target - root.Val
		findTargetAux(root.Right, numberTable, target)
	}
}
