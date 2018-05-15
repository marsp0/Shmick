package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
	var currentNode = head
	for {
		if currentNode == nil {
			break
		}
		if currentNode.Next == nil {
			break
		}
		if currentNode.Val == currentNode.Next.Val {
			currentNode.Next = currentNode.Next.Next
			continue
		}
		currentNode = currentNode.Next
	}
	return head
}
