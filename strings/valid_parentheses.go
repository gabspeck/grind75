package strings

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
//	Open brackets must be closed by the same type of brackets.
//	Open brackets must be closed in the correct order.
//	Every close bracket has a corresponding open bracket of the same type.
//
// Example 1:
//
// Input: s = "()"
//
// Output: true
//
// Example 2:
//
// Input: s = "()[]{}"
//
// Output: true
//
// Example 3:
//
// Input: s = "(]"
//
// Output: false
//
// Example 4:
//
// Input: s = "([])"
//
// Output: true
//
// Constraints:
//
//	1 <= s.length <= 104
//	s consists of parentheses only '()[]{}'.
type ValidParenthesesFunc func(s string) bool

func ValidParentheses(s string) bool {

	openers := ""

	pairs := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, c := range []rune(s) {
		if opener, isCloser := pairs[c]; isCloser {
			if len(openers) > 0 {
				if rune(openers[len(openers)-1]) != opener {
					return false
				}
				openers = openers[0 : len(openers)-1]
			} else {
				return false
			}
		} else {
			openers += string(c)
		}
	}

	return openers == ""

}
