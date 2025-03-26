package arrays

// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
type MajorityElementFunc func([]int) int

// MajorityElementReference is an implementation of the Boyer-Moore majority vote algorithm
func MajorityElementReference(nums []int) int {
	var candidate int
	count := 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}

	return candidate
}
