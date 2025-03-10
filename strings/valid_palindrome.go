package strings

import (
	"regexp"
	"strings"
	"unicode"
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

// IsPalindromeAttempt2 is a second attempt based on feedback from Claude
func IsPalindromeAttempt2(s string) bool {

	i := 0
	j := len(s) - 1

	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
	s = strings.ToLower(s)

	for i < len(s) && j >= 0 {
		charFromStart := rune(s[i])
		charFromEnd := rune(s[j])
		for !strings.ContainsRune(alphabet, charFromStart) {
			i++
			if i >= len(s) {
				break
			}
			charFromStart = rune(s[i])
		}
		for !strings.ContainsRune(alphabet, charFromEnd) {
			j--
			if j < 0 {
				break
			}
			charFromEnd = rune(s[j])
		}
		if strings.ContainsRune(alphabet, charFromStart) && strings.ContainsRune(alphabet, charFromEnd) && charFromStart != charFromEnd {
			return false
		}
		i++
		j--
	}

	return true

}

// IsPalindromeReference is a reference solution generated by Claude
func IsPalindromeReference(s string) bool {
	// Convert string to runes for proper Unicode handling
	runes := []rune(strings.ToLower(s))

	var cleaned []rune
	// Keep only alphanumeric characters
	for _, r := range runes {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}

	// Check if the cleaned string is a palindrome
	left, right := 0, len(cleaned)-1
	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}
