package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// https://leetcode.com/problems/merge-two-sorted-lists/description/
	var (
		currentLeft  *ListNode
		currentRight *ListNode
		result       *ListNode
		head         *ListNode
	)

	if l1 == nil || l2 == nil {
		if l1 == nil {
			return l2
		}
		return l1
	}

	if l1.Val < l2.Val {
		result = l1
		head = l1
		currentLeft = l1.Next
		currentRight = l2
	} else {
		result = l2
		head = l2
		currentRight = l2.Next
		currentLeft = l1
	}
	for {
		if currentRight != nil && currentLeft != nil {
			if currentLeft.Val < currentRight.Val {
				result.Next = currentLeft
				result = result.Next
				currentLeft = currentLeft.Next
			} else {
				result.Next = currentRight
				result = result.Next
				currentRight = currentRight.Next
			}
		} else {
			if currentLeft != nil {
				result.Next = currentLeft
				break
			}
			if currentRight != nil {
				result.Next = currentRight
				break
			}
		}
	}
	return head
}
