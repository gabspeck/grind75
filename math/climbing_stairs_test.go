package math_test

import (
	"grind75/math"
	"testing"
	"time"
)

import (
	"fmt"
)

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{45, 1836311903},
	}
	impls := map[string]math.ClimbStairsFunc{
		"ClimbStairsNaive":              math.ClimbStairsNaive,
		"ClimbStairsAttemptMemo":        math.ClimbStairsAttemptMemo,
		"ClimbStairsIterativeReference": math.ClimbStairsIterativeReference,
		"ClimbStairsUsingBinet":         math.ClimbStairsUsingBinet,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (n=\"%v\")", name, testCase.n), func(t *testing.T) {
				timeout := time.After(time.Millisecond)
				done := make(chan bool)
				go func() {
					actual := impl(testCase.n)
					if actual != testCase.expected {
						t.Errorf("wanted %v; got %v", testCase.expected, actual)
					}
					done <- true
				}()
				select {
				case <-timeout:
					t.Fatal("timeout exceeded!")
				case <-done:
				}
			})
		}
	}
}
