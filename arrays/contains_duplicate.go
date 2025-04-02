package arrays

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
type ContainsDuplicateFunc func([]int) bool

// ContainsDuplicateAttempt uses a hash map to flag the first time it sees a number and returns once it's seen for a second time
func ContainsDuplicateAttempt(nums []int) bool {
	seen := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if seen[nums[i]] {
			return true
		}
		seen[nums[i]] = true
	}
	return false
}
