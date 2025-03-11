package strings_test

import (
	"grind75/strings"
	"testing"
)

import (
	"fmt"
)

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"ac", "ab", false},
	}
	impls := map[string]strings.IsAnagramFunc{
		"IsAnagramAttempt":  strings.IsAnagramAttempt,
		"IsAnagramAttempt2": strings.IsAnagramAttempt2,
		"IsAnagramAttempt3": strings.IsAnagramAttempt3,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (s1=\"%s\", s2=\"%s\")", name, testCase.s1, testCase.s2), func(t *testing.T) {
				actual := impl(testCase.s1, testCase.s2)
				if actual != testCase.expected {
					t.Errorf("wanted %v; got %v", testCase.expected, actual)
				}
			})
		}
	}
}
