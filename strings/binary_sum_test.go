package strings_test

import (
	"grind75/strings"
	"testing"
)

import (
	"fmt"
)

func TestAddBinary(t *testing.T) {
	cases := []struct {
		a        string
		b        string
		expected string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011", "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"},
	}
	impls := map[string]strings.AddBinaryFunc{
		"AddBinaryAttempt":  strings.AddBinaryAttempt,
		"AddBinaryAttempt2": strings.AddBinaryAttempt2,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (a=\"%s\", b=\"%s\")", name, testCase.a, testCase.b), func(t *testing.T) {
				actual := impl(testCase.a, testCase.b)
				if actual != testCase.expected {
					t.Errorf("wanted \"%v\"; got \"%v\"", testCase.expected, actual)
				}
			})
		}
	}
}
