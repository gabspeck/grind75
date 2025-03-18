package linked_lists_test

import (
	"fmt"
	"grind75/linked_lists"
	"testing"
)

func TestHasCycle(t *testing.T) {
	cases := []struct {
		list     *linked_lists.ListNode
		expected bool
	}{
		{
			list: func() *linked_lists.ListNode {
				head := linked_lists.NewLinkedList(3, 2, 0, -4)
				head.Tail().Next = head.Next
				return head
			}(),
			expected: true,
		},
		{
			list: func() *linked_lists.ListNode {
				head := linked_lists.NewLinkedList(1, 2)
				head.Tail().Next = head
				return head
			}(),
			expected: true,
		},
		{
			list:     linked_lists.NewLinkedList(1),
			expected: false,
		},
	}
	impls := map[string]linked_lists.HasCycleFunc{
		"HasCycleAttempt": linked_lists.HasCycleAttempt,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (list=%v)", name, testCase.list), func(t *testing.T) {
				actual := impl(testCase.list)
				if actual != testCase.expected {
					t.Errorf("wanted %v, got %v", testCase.expected, actual)
				}
			})
		}
	}
}
