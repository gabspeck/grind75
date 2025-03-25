package linked_lists_test

import (
	"fmt"
	"grind75/linked_lists"
	"testing"
)

func TestMergeTwoListsAttempt(t *testing.T) {
	cases := []struct {
		list1    *linked_lists.ListNode
		list2    *linked_lists.ListNode
		expected *linked_lists.ListNode
	}{
		{
			list1:    linked_lists.New(1, 2, 4),
			list2:    linked_lists.New(1, 3, 4),
			expected: linked_lists.New(1, 1, 2, 3, 4, 4),
		},
		{
			list1:    nil,
			list2:    nil,
			expected: nil,
		},
		{
			list1:    nil,
			list2:    linked_lists.New(0),
			expected: linked_lists.New(0),
		},
	}
	impls := map[string]linked_lists.MergeTwoListsFunc{
		"MergeTwoListsAttempt":   linked_lists.MergeTwoListsAttempt1,
		"MergeTwoListsReference": linked_lists.MergeTwoListsReference,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (list1=%v, list2=%v)", name, testCase.list1, testCase.list2), func(t *testing.T) {
				actual := impl(testCase.list1, testCase.list2)
				if actual.String() != testCase.expected.String() {
					t.Errorf("wanted %v, got %v", testCase.expected, actual)
				}
			})
		}
	}
}
