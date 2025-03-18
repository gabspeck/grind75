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

func (head *ListNode) String() string {
	builder := strings.Builder{}
	builder.WriteRune('[')

	for el := head; el != nil; el = el.Next {
		builder.WriteString(fmt.Sprintf("%d", el.Val))
		if el.Next != nil {
			builder.WriteRune(',')
		}
	}
	builder.WriteRune(']')

	return builder.String()
}

func (head *ListNode) Tail() *ListNode {
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	return tail
}
