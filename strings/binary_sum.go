package strings

// Given two binary strings a and b, return their sum as a binary string.
//
// Example 1:
//
// Input: a = "11", b = "1"
// Output: "100"
//
// Example 2:
//
// Input: a = "1010", b = "1011"
// Output: "10101"
//
// Constraints:
//
//	1 <= a.length, b.length <= 104
//	a and b consist only of '0' or '1' characters.
//	Each string does not contain leading zeros except for the zero itself.
type AddBinaryFunc func(a, b string) string

func parseBinary(s string) int {
	n := 0
	length := len(s)
	for i, digit := range s {
		power := length - i - 1
		if digit == '1' {
			n += 1 << power
		}
	}
	return n
}

// AddBinaryAttempt first converts a and b to ints, which works but is limited to numbers that can fit in 64 bits
func AddBinaryAttempt(a, b string) string {

	quotient := parseBinary(a) + parseBinary(b)
	resultString := ""

	if quotient == 0 {
		return "0"
	}
	for quotient > 0 {
		if quotient%2 == 0 {
			resultString = "0" + resultString
		} else {
			resultString = "1" + resultString
		}
		quotient /= 2
	}

	return resultString

}

// AddBinaryAttempt2 parses the digits right to left and reproduces the manual sum process
func AddBinaryAttempt2(a, b string) string {
	aPtr, bPtr := len(a)-1, len(b)-1

	carry := 0
	result := ""
	for aPtr >= 0 || bPtr >= 0 {
		aBit, bBit := 0, 0
		if aPtr > -1 {
			if a[aPtr] == '1' {
				aBit = 1
			}
			aPtr--
		}
		if bPtr > -1 {
			if b[bPtr] == '1' {
				bBit = 1
			}
			bPtr--
		}
		sum := aBit + bBit + carry
		bit := "0"
		if sum%2 != 0 {
			bit = "1"
		}
		result = bit + result
		carry = sum / 2
	}
	if carry > 0 {
		result = "1" + result
	}
	return result
}
