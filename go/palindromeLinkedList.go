package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeLL(head *ListNode) bool {
	// https://leetcode.com/problems/palindrome-linked-list/description/
	var (
		array   = []int{}
		current *ListNode
		i       = 0
		j       = 0
	)
	current = head
	for current != nil {
		array = append(array, current.Val)
		current = current.Next
	}
	j = len(array) - 1
	for i < j {
		if array[i] != array[j] {
			return false
		}
		i++
		j--
	}
	return true
}
