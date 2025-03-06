package arrays_test

import (
	"fmt"
	"grind75/arrays"
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		}, {
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}
	impls := map[string]arrays.TwoSumFunc{
		"TwoSumBruteForce": arrays.TwoSumBruteForce,
		"TwoSumOptimized":  arrays.TwoSumOptimized,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (nums=%v, target=%v)", name, testCase.nums, testCase.target), func(t *testing.T) {
				actual := impl(testCase.nums, testCase.target)
				if slices.Compare(actual, testCase.expected) != 0 {
					t.Errorf("got %v, wanted %v", actual, testCase.expected)
				}
			})
		}
	}
}
