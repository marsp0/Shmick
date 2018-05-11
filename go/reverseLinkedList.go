package main

// ListNode - implementation of a node in Linked List
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// https://leetcode.com/problems/reverse-linked-list/description/
	var (
		current = head
		prev    *ListNode
	)
	for {
		if current != nil {
			var temp = current.Next
			current.Next = prev
			prev = current
			current = temp
		} else {
			current = prev
			break
		}
	}
	return current
}
