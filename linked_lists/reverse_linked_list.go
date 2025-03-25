package linked_lists

// Given the head of a singly linked list, reverse the list, and return the reversed list.
type ReverseListFunc func(*ListNode) *ListNode

// ReverseListAttempt is a first attempt at solving the problem
func ReverseListAttempt(head *ListNode) *ListNode {
	nodeStack := make([]*ListNode, 0)

	for tail := head; tail != nil; tail = tail.Next {
		nodeStack = append(nodeStack, tail)
	}
	if len(nodeStack) == 0 {
		return nil
	}
	newHead := nodeStack[len(nodeStack)-1]
	if len(nodeStack) < 2 {
		return newHead
	}
	for i := len(nodeStack) - 1; i >= 0; i-- {
		if i > 0 {
			nodeStack[i].Next = nodeStack[i-1]
		} else {
			nodeStack[i].Next = nil
		}
	}
	return newHead
}

// ReverseListAttempt2 is a new attempt with O(1) space complexity and that traverses the list a single time instead of twice
func ReverseListAttempt2(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	prev := head
	cur := head.Next
	prev.Next = nil

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
