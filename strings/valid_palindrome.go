package strings

import (
	"regexp"
	"strings"
)

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//
// Example 1:
//
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
//
// Example 2:
//
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
//
// Example 3:
//
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.
//
// Constraints:
//
//	1 <= s.length <= 2 * 105
//	s consists only of printable ASCII characters.
type IsPalindromeFunc func(s string) bool

func IsPalindromeAttempt(s string) bool {

	s = strings.ToLower(s)
	nonLettersRegex := regexp.MustCompile("[^a-z0-9]")
	s = nonLettersRegex.ReplaceAllString(s, "")

	letterStack := make([]rune, 0)

	for i, r := range s {
		if i < len(s)/2 {
			letterStack = append(letterStack, r)
		} else if len(s)%2 != 0 && i == len(s)/2 {
			continue
		} else {
			topLetter := letterStack[len(letterStack)-1]
			if r != topLetter {
				return false
			}
			letterStack = letterStack[:len(letterStack)-1]
		}
	}

	return len(letterStack) == 0
}
