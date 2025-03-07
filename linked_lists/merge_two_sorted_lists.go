package linked_lists

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(items ...int) *ListNode {
	if len(items) < 1 {
		return nil
	}
	head := &ListNode{items[0], nil}
	tail := head
	for i := 1; i < len(items); i++ {
		tail.Next = &ListNode{items[i], nil}
		tail = tail.Next
	}
	return head
}

func (ln *ListNode) String() string {
	builder := strings.Builder{}
	builder.WriteRune('[')
	for el := ln; el != nil; el = el.Next {
		builder.WriteString(fmt.Sprintf("%d", el.Val))
		if el.Next != nil {
			builder.WriteRune(',')
		}
	}
	builder.WriteRune(']')

	return builder.String()
}

// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.
type MergeTwoListsFunc func(list1 *ListNode, list2 *ListNode) *ListNode

// MergeTwoListsAttempt1 creates copies instead of reusing the nodes, which is probably wrong lol
func MergeTwoListsAttempt1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := &ListNode{}
	if list1.Val < list2.Val {
		head.Val = list1.Val
		list1 = list1.Next
	} else {
		head.Val = list2.Val
		list2 = list2.Next
	}
	tail := head

	for {
		if list1 != nil && (list2 == nil || list1.Val < list2.Val) {
			tail.Next = &ListNode{list1.Val, nil}
			list1 = list1.Next
		} else if list2 != nil {
			tail.Next = &ListNode{list2.Val, nil}
			list2 = list2.Next
		}
		tail = tail.Next
		if tail == nil {
			break
		}
	}

	return head
}

// MergeTwoListsReference reuses the existing nodes
func MergeTwoListsReference(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := &ListNode{}
	tail := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return head.Next
}
