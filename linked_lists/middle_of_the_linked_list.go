package linked_lists

// Given the head of a singly linked list, return the middle node of the linked list.
//
// If there are two middle nodes, return the second middle node.
type MiddleOfTheLinkedListFunc func(head *ListNode) *ListNode

func MiddleOfTheLinkedListAttempt(head *ListNode) *ListNode {
	listLen := 0
	for elem := head; elem != nil; elem = elem.Next {
		listLen++
	}
	middle := listLen / 2
	middleNode := head
	for i := 0; i < middle; i++ {
		middleNode = middleNode.Next
	}

	return middleNode
}

// MiddleOfTheLinkedListReference is the Claude-generated reference solution
func MiddleOfTheLinkedListReference(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
