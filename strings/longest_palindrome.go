package strings

// Given a string s which consists of lowercase or uppercase letters, return the length of the longest
//
// that can be built with those letters.
//
// Letters are case sensitive, for example, "Aa" is not considered a palindrome.
//
// Example 1:
//
// Input: s = "abccccdd"
// Output: 7
// Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.
//
// Example 2:
//
// Input: s = "a"
// Output: 1
// Explanation: The longest palindrome that can be built is "a", whose length is 1.
//
// Constraints:
//
//	1 <= s.length <= 2000
//	s consists of lowercase and/or uppercase English letters only.
type LongestPalindromeFunc func(s string) int

func LongestPalindromeAttempt(s string) int {
	counts := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		counts[s[i]] += 1
	}

	solution := 0
	foundOdd := false

	for _, v := range counts {
		if v%2 == 0 {
			solution += v
		} else {
			foundOdd = true
			solution += v - 1
		}
	}

	if foundOdd {
		solution++
	}

	return solution

}
