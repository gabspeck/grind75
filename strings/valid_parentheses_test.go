package strings_test

import (
	"grind75/strings"
	"testing"
)

import (
	"fmt"
)

func TestValidParentheses(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{
			s:        "()",
			expected: true,
		},
		{
			s:        "()[]{}",
			expected: true,
		}, {
			s:        "(]",
			expected: false,
		}, {
			s:        "([])",
			expected: true,
		}, {
			s:        "[",
			expected: false,
		}, {
			s:        "]",
			expected: false,
		}, {
			s:        "((",
			expected: false,
		}, {
			s:        ")(){}",
			expected: false,
		},
	}
	impls := map[string]strings.ValidParenthesesFunc{
		"IsValid": strings.ValidParentheses,
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
