package arrays_test

import (
	"fmt"
	"grind75/arrays"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}
	impls := map[string]arrays.BinarySearchFunc{
		"LinearSearch":                 arrays.LinearSearch,
		"BinarySearchAttemptRecursive": arrays.BinarySearchAttemptRecursive,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (arr=%v, target=%v)", name, testCase.arr, testCase.target), func(t *testing.T) {
				actual := impl(testCase.arr, testCase.target)
				if actual != testCase.expected {
					t.Errorf("got %v, wanted %v", actual, testCase.expected)
				}
			})
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	// Create test arrays of different sizes
	sizes := []int{1_000_000_000}

	// Define implementations to benchmark
	impls := map[string]arrays.BinarySearchFunc{
		"LinearSearch":                 arrays.LinearSearch,
		"BinarySearchAttemptRecursive": arrays.BinarySearchAttemptRecursive,
		"BinarySearchReference":        arrays.BinarySearchReference,
	}

	for _, size := range sizes {
		// Create a sorted array of the specified size
		arr := make([]int, size)
		for i := range arr {
			arr[i] = i
		}

		// Test different scenarios: beginning, middle, end, not found
		scenarios := []struct {
			name   string
			target int
		}{
			{"Beginning", 0},
			{"Middle", size / 2},
			{"End", size - 1},
			{"NotFound", size + 100},
		}

		// Benchmark each implementation with each scenario
		for implName, impl := range impls {
			for _, sc := range scenarios {
				b.Run(fmt.Sprintf("%s/Size=%d/%s", implName, size, sc.name), func(b *testing.B) {
					// Reset timer to exclude setup time
					b.ResetTimer()

					// Run the benchmark
					for i := 0; i < b.N; i++ {
						impl(arr, sc.target)
					}
				})
			}
		}
	}
}
