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
		{".,", true},
	}
	impls := map[string]strings.IsPalindromeFunc{
		"IsPalindromeAttempt":   strings.IsPalindromeAttempt,
		"IsPalindromeAttempt2":  strings.IsPalindromeAttempt2,
		"IsPalindromeReference": strings.IsPalindromeReference,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (string=\"%s\")", name, testCase.s), func(t *testing.T) {
				actual := impl(testCase.s)
				if actual != testCase.expected {
					t.Errorf("wanted %v; got %v", testCase.expected, actual)
				}
			})
		}
	}
}
