package arrays_test

import (
	"fmt"
	"grind75/arrays"
	"testing"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}},
		},
		{
			[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}
	impls := map[string]arrays.InsertIntervalFunc{
		"InsertIntervalReference": arrays.InsertIntervalReference,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			intervals := arrays.Copy2DSlice(testCase.intervals)
			t.Run(fmt.Sprintf("%s (intervals=%v, newInterval=%v)", name, intervals, testCase.newInterval), func(t *testing.T) {
				actual := impl(intervals, testCase.newInterval)
				if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", testCase.expected) {
					t.Errorf("got %v, wanted %v", actual, testCase.expected)
				}
			})
		}
	}
}
