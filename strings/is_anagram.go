package strings

import (
	"slices"
	"strings"
)

type IsAnagramFunc func(string, string) bool

// IsAnagramAttempt is a naïve approach that equalizes both strings and compares them. Works for any Unicode string though
func IsAnagramAttempt(s, t string) bool {
	runes1 := []rune(s)
	runes2 := []rune(t)

	slices.Sort(runes1)
	slices.Sort(runes2)

	return slices.Compare(runes1, runes2) == 0

}

// IsAnagramAttempt2 implements a frequency counter but uses an index lookup every iteration, which is not optimal
func IsAnagramAttempt2(s, t string) bool {

	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	if len(s) != len(t) {
		return false
	}

	counts := make([]int, len(alphabet))

	for i := 0; i < len(s); i++ {
		counts[strings.IndexRune(alphabet, rune(s[i]))] += 1
	}
	for i := 0; i < len(t); i++ {
		counts[strings.IndexRune(alphabet, rune(t[i]))] -= 1
	}

	for i := 0; i < len(counts); i++ {
		if counts[i] != 0 {
			return false
		}
	}

	return true
}

// IsAnagramAttempt3 takes advantage of the sequential char codes for the latin alphabet, so only works for lowercase ASCII strings
func IsAnagramAttempt3(s, t string) bool {

	const alphabetIndexOffset = 'a'

	if len(s) != len(t) {
		return false
	}

	counts := make([]int, 26)

	for i := 0; i < len(s); i++ {
		counts[s[i]-alphabetIndexOffset]++
		counts[t[i]-alphabetIndexOffset]--
	}

	for i := 0; i < len(counts); i++ {
		if counts[i] != 0 {
			return false
		}
	}

	return true
}
