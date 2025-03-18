package strings_test

import (
	"grind75/strings"
	"testing"
)

import (
	"fmt"
)

func TestCanConstruct(t *testing.T) {
	cases := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}
	impls := map[string]strings.CanConstructFunc{
		"CanConstructAttempt": strings.CanConstructAttempt,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (ransomNote=\"%s\", magazine=\"%s\")", name, testCase.ransomNote, testCase.magazine), func(t *testing.T) {
				actual := impl(testCase.ransomNote, testCase.magazine)
				if actual != testCase.expected {
					t.Errorf("wanted %v; got %v", testCase.expected, actual)
				}
			})
		}
	}
}
