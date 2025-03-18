package strings

// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
//
// Each letter in magazine can only be used once in ransomNote.
//
// Example 1:
//
// Input: ransomNote = "a", magazine = "b"
// Output: false
//
// Example 2:
//
// Input: ransomNote = "aa", magazine = "ab"
// Output: false
//
// Example 3:
//
// Input: ransomNote = "aa", magazine = "aab"
// Output: true
//
// Constraints:
//
//	1 <= ransomNote.length, magazine.length <= 105
//	ransomNote and magazine consist of lowercase English letters.
type CanConstructFunc func(ransomNote, magazine string) bool

func CanConstructAttempt(ransomNote, magazine string) bool {

	const alphabetOffset = 'a'
	magazineFreqs := make([]int, 26)

	for _, r := range magazine {
		magazineFreqs[r-alphabetOffset]++
	}

	for _, r := range ransomNote {
		magazineFreqs[r-alphabetOffset]--
		if magazineFreqs[r-alphabetOffset] < 0 {
			return false
		}
	}

	return true

}
