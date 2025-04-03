package arrays

// Given an integer array nums, find the with the largest sum, and return its sum.
type MaxSubArrayFunc func([]int) int

func MaxSubArrayAttempt(nums []int) int {
	sum := nums[0]
	maxSum := sum
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		sum = max(n, sum+n)
		maxSum = max(sum, maxSum)
	}

	return maxSum
}
