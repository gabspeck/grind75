package strings_test

import (
	"grind75/strings"
	"testing"
)

import (
	"fmt"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
	}
	impls := map[string]strings.IsPalindromeFunc{
		"IsPalindrome": strings.IsPalindromeAttempt,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (string=\"%s\")", name, testCase.s), func(t *testing.T) {
				actual := impl(testCase.s)
				if actual != testCase.expected {
					t.Errorf("wanted %v; got %v", actual, testCase.expected)
				}
			})
		}
	}
}
