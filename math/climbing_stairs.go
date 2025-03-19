package math

import "math"

// You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Example 1:
//
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
//
// Example 2:
//
// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
//
// Constraints:
//
//	1 <= n <= 45
type ClimbStairsFunc func(n int) int

func ClimbStairsNaive(n int) int {
	if n < 3 {
		return n
	}
	return ClimbStairsNaive(n-1) + ClimbStairsNaive(n-2)
}

// ClimbStairsAttemptMemo is an attempt at memoizing the results of the recursive calculation
func ClimbStairsAttemptMemo(n int) int {
	memo := make(map[int]int)

	var impl ClimbStairsFunc
	impl = func(n int) int {
		if n < 3 {
			return n
		}

		if memoValue, isMemoized := memo[n]; isMemoized {
			return memoValue
		}

		memo[n] = impl(n-2) + impl(n-1)

		return memo[n]
	}
	return impl(n)
}

// ClimbStairsIterativeReference is an iterative reference solution generated by Claude
func ClimbStairsIterativeReference(n int) int {
	if n < 3 {
		return n
	}

	first, second := 1, 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}

// ClimbStairsUsingBinet calculates the number of ways to climb n stairs
// using Binet's mathematical formula for Fibonacci numbers (Claude-generated)
func ClimbStairsUsingBinet(n int) int {
	// Handle edge cases separately
	if n < 0 {
		return 0 // Invalid input
	}

	// Calculate golden ratio φ = (1 + √5) / 2 and its complement
	phi := (1 + math.Sqrt(5)) / 2
	psi := (1 - math.Sqrt(5)) / 2

	// Apply Binet's formula for F(n+1) to get the number of ways to climb n stairs
	result := (math.Pow(phi, float64(n+1)) - math.Pow(psi, float64(n+1))) / math.Sqrt(5)

	// Round to nearest integer to handle floating-point precision
	return int(math.Round(result))
}
