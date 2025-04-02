package arrays_test

import (
	"fmt"
	"grind75/arrays"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	cases := []struct {
		arr      []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	impls := map[string]arrays.ContainsDuplicateFunc{
		"ContainsDuplicateAttempt": arrays.ContainsDuplicateAttempt,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (arr=%v)", name, testCase.arr), func(t *testing.T) {
				actual := impl(testCase.arr)
				if actual != testCase.expected {
					t.Errorf("got %v, wanted %v", actual, testCase.expected)
				}
			})
		}
	}
}
