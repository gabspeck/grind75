package linked_lists_test

import (
	"fmt"
	"grind75/linked_lists"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	cases := []struct {
		head     *linked_lists.ListNode
		expected *linked_lists.ListNode
	}{
		{
			head:     linked_lists.New(1, 2, 3, 4, 5),
			expected: linked_lists.New(5, 4, 3, 2, 1),
		},
		{
			head:     linked_lists.New(1, 2),
			expected: linked_lists.New(2, 1),
		},
		{
			head:     nil,
			expected: nil,
		},
	}
	impls := map[string]linked_lists.ReverseListFunc{
		"ReverseListAttempt":  linked_lists.ReverseListAttempt,
		"ReverseListAttempt2": linked_lists.ReverseListAttempt2,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (head=%v)", name, testCase.head), func(t *testing.T) {
				actual := impl(linked_lists.Copy(testCase.head))
				if actual.String() != testCase.expected.String() {
					t.Errorf("wanted %v, got %v", testCase.expected, actual)
				}
			})
		}
	}
}
