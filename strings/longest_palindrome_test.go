package strings_test

import (
	"grind75/strings"
	"testing"
)

import (
	"fmt"
)

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{"abccccdd", 7},
		{"a", 1},
	}
	impls := map[string]strings.LongestPalindromeFunc{
		"LongestPalindromeAttempt": strings.LongestPalindromeAttempt,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (s1=\"%s\")", name, testCase.s), func(t *testing.T) {
				actual := impl(testCase.s)
				if actual != testCase.expected {
					t.Errorf("wanted %v; got %v", testCase.expected, actual)
				}
			})
		}
	}
}
