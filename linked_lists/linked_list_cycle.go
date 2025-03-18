package linked_lists

// Given head, the head of a linked list, determine if the linked list has a cycle in it.
//
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
//
// Return true if there is a cycle in the linked list. Otherwise, return false.
type HasCycleFunc func(head *ListNode) bool

var seen = make(map[*ListNode]bool)

// HasCycleAttempt is a first attempt that uses at most O(N) space
func HasCycleAttempt(head *ListNode) bool {
	node := head
	for node != nil {
		if seen[node] {
			return true
		}
		seen[node] = true
		node = node.Next
	}

	return false
}

// HasCycleReferenceAttempt is an attempt based on advice given by Claude.
// It implements Floyd's Cycle-Finding Algorithm
func HasCycleReferenceAttempt(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false
}
