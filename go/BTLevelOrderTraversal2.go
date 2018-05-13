package main

func levelOrderBottom(root *TreeNode) (result [][]int) {
	// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
	var (
		queue    = [][]*TreeNode{}
		levelMap = map[*TreeNode]int{}
		i        int
		j        int
	)
	queue = append(queue, []*TreeNode{root, nil})
	levelMap[root] = 0
	for len(queue) > 0 {
		var (
			nodes       = queue[0]
			currentNode = nodes[0]
			parentNode  = nodes[1]
		)
		queue = queue[1:]
		if currentNode != nil {
			if parentNode != nil {
				if levelMap[parentNode]+1 >= len(result) {
					result = append(result, []int{currentNode.Val})
				} else {
					result[levelMap[parentNode]+1] = append(result[levelMap[parentNode]+1], currentNode.Val)
				}
				levelMap[currentNode] = levelMap[parentNode] + 1
			} else {
				result = append(result, []int{currentNode.Val})
			}
			queue = append(queue, []*TreeNode{currentNode.Left, currentNode})
			queue = append(queue, []*TreeNode{currentNode.Right, currentNode})
		}
	}
	j = len(result) - 1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}
