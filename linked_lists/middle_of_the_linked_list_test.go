package linked_lists_test

import (
	"fmt"
	"grind75/linked_lists"
	"testing"
)

func TestMiddleOfTheLinkedList(t *testing.T) {
	cases := []struct {
		list     *linked_lists.ListNode
		expected int
	}{
		{
			list:     linked_lists.New(1, 2, 3, 4, 5),
			expected: 3,
		},
		{
			list:     linked_lists.New(1, 2, 3, 4, 5, 6),
			expected: 4,
		},
	}
	impls := map[string]linked_lists.MiddleOfTheLinkedListFunc{
		"MiddleOfTheLinkedListAttempt":   linked_lists.MiddleOfTheLinkedListAttempt,
		"MiddleOfTheLinkedListReference": linked_lists.MiddleOfTheLinkedListReference,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (list=%v)", name, testCase.list), func(t *testing.T) {
				actual := impl(testCase.list)
				if actual.Val != testCase.expected {
					t.Errorf("wanted %v, got %v", testCase.expected, actual)
				}
			})
		}
	}
}
func BenchmarkMiddleOfTheLinkedList(b *testing.B) {
	// Define different sizes of linked lists for benchmarking
	sizes := []int{10, 100, 1000, 10000}

	// Get implementations to benchmark
	impls := map[string]linked_lists.MiddleOfTheLinkedListFunc{
		"Attempt":   linked_lists.MiddleOfTheLinkedListAttempt,
		"Reference": linked_lists.MiddleOfTheLinkedListReference,
	}

	// Run benchmarks for each implementation with different list sizes
	for name, impl := range impls {
		for _, size := range sizes {
			// Test with both even and odd list sizes
			for _, adjustment := range []int{0, 1} {
				actualSize := size + adjustment
				b.Run(fmt.Sprintf("%s-Size%d", name, actualSize), func(b *testing.B) {
					// Create a linked list of the given size
					list := createLinkedList(actualSize)

					// Reset the timer before the benchmark loop
					b.ResetTimer()

					// Run the implementation b.N times
					for i := 0; i < b.N; i++ {
						_ = impl(list)
					}
				})
			}
		}
	}
}

// Helper function to create a linked list of specified size
func createLinkedList(size int) *linked_lists.ListNode {
	var head *linked_lists.ListNode
	for i := size; i > 0; i-- {
		head = &linked_lists.ListNode{Val: i, Next: head}
	}
	return head
}
