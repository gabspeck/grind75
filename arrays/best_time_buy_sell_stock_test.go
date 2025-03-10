package arrays_test

import (
	"fmt"
	"grind75/arrays"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	impls := map[string]arrays.MaxProfitFunc{
		"MaxProfitAttempt":   arrays.MaxProfitAttempt,
		"MaxProfitReference": arrays.MaxProfitReference,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (prices=%v)", name, testCase.prices), func(t *testing.T) {
				actual := impl(testCase.prices)
				if actual != testCase.expected {
					t.Errorf("got %v, wanted %v", actual, testCase.expected)
				}
			})
		}
	}
}
