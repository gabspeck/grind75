package arrays_test

import (
	"fmt"
	"grind75/arrays"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}
	impls := map[string]arrays.MaxSubArrayFunc{
		"MaxSubArrayAttempt": arrays.MaxSubArrayAttempt,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (nums=%v)", name, testCase.nums), func(t *testing.T) {
				actual := impl(testCase.nums)
				if actual != testCase.expected {
					t.Errorf("got %v, wanted %v", actual, testCase.expected)
				}
			})
		}
	}
}
