package arrays

// Copy2DSlice creates a deep copy of a 2D integer slice.
// This is necessary because Go's copy() function only performs shallow copies.
func Copy2DSlice(original [][]int) [][]int {
	imageCopy := make([][]int, len(original))
	for i := range original {
		imageCopy[i] = make([]int, len(original[i]))
		for j := range original[i] {
			imageCopy[i][j] = original[i][j]
		}
	}
	return imageCopy
}
