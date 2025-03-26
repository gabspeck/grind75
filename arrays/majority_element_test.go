package arrays_test

import (
	"fmt"
	"grind75/arrays"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	impls := map[string]arrays.MajorityElementFunc{
		"MajorityElementReference": arrays.MajorityElementReference,
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
