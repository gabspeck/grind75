package arrays

type TwoSumFunc func(nums []int, target int) []int

func TwoSumBruteForce(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func TwoSumOptimized(nums []int, target int) []int {

	/*
		The main inefficiency is that we're repeatedly checking combinations we've already considered. For example,
		if we've checked nums[2] + nums[5], we don't need to check nums[5] + nums[2] later.
	*/

	// O(N) space complexity - memory-performance tradeoff
	seenValuesIndices := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if val, exists := seenValuesIndices[complement]; exists {
			return []int{val, i}
		}
		seenValuesIndices[nums[i]] = i
	}

	return []int{}
}
