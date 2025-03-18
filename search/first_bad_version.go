package search

func isBadVersion(version int) bool {
	return version > 3
}

func FirstBadVersionAttempt(n int) int {
	low := 1
	high := n - 1

	for low <= high {
		mid := low + (high-low)/2
		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
