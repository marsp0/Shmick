package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// https://leetcode.com/problems/remove-linked-list-elements/description/
	var (
		current = head
		prev    = head
	)
	if head != nil && head.Val == val {
		return removeElements(head.Next, val)
	}
	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
			current = current.Next
			continue
		}
		prev = current
		current = current.Next

	}
	return head
}
