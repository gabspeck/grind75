package arrays_test

import (
	"fmt"
	"grind75/arrays"
	"testing"
)

func TestFloodFill(t *testing.T) {
	cases := []struct {
		image    [][]int
		sr       int
		sc       int
		color    int
		expected [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{[][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0, [][]int{{0, 0, 0}, {0, 0, 0}}},
	}
	impls := map[string]arrays.FloodFillFunc{
		"FloodFillAttemptRecursive":          arrays.FloodFillAttemptRecursive,
		"FloodFillAttemptIterative":          arrays.FloodFillAttemptIterative,
		"FloodFillAttemptIterativeOptimized": arrays.FloodFillAttemptIterativeOptimized,
		"FloodFillReference":                 arrays.FloodFillReference,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			image := makeImageCopy(testCase.image)
			t.Run(fmt.Sprintf("%s (image=%v, sr=%v, sc=%v, color=%v)", name, image, testCase.sr, testCase.sc, testCase.color), func(t *testing.T) {
				actual := impl(image, testCase.sr, testCase.sc, testCase.color)
				if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", testCase.expected) {
					t.Errorf("got %v, wanted %v", actual, testCase.expected)
				}
			})
		}
	}
}

// BenchmarkFloodFill benchmarks different flood fill implementations
// with various image sizes to compare performance characteristics.
func BenchmarkFloodFill(b *testing.B) {
	// Define test cases of varying sizes and complexity
	cases := []struct {
		name  string
		image [][]int
		sr    int
		sc    int
		color int
	}{
		{"small_3x3", [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2},
		{"small_2x3", [][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0},
		{"medium_10x10", generateFilledImage(10, 10, 1), 0, 0, 2},
		{"large_50x50", generateFilledImage(50, 50, 1), 0, 0, 2},
		{"mixed_20x20", generateMixedImage(20, 20), 10, 10, 5}, // More realistic image with mixed colors
	}

	// The implementations to test
	impls := map[string]arrays.FloodFillFunc{
		"Recursive":           arrays.FloodFillAttemptRecursive,
		"Iterative":           arrays.FloodFillAttemptIterative,
		"Optimized Iterative": arrays.FloodFillAttemptIterativeOptimized,
		"Reference":           arrays.FloodFillReference,
	}

	// Run benchmarks for each implementation and test case
	for name, impl := range impls {
		for _, tc := range cases {
			b.Run(fmt.Sprintf("%s/%s", name, tc.name), func(b *testing.B) {

				// Reset the timer before the benchmark loop
				b.ResetTimer()

				// Run the benchmark b.N times
				for i := 0; i < b.N; i++ {
					// Create a fresh copy of the image for each iteration
					imageCopy := makeImageCopy(tc.image)

					// Call the implementation being benchmarked
					_ = impl(imageCopy, tc.sr, tc.sc, tc.color)
				}
			})
		}
	}
}

// makeImageCopy creates a deep copy of a 2D integer slice.
// This is necessary because Go's copy() function only performs shallow copies.
func makeImageCopy(original [][]int) [][]int {
	imageCopy := make([][]int, len(original))
	for i := range original {
		imageCopy[i] = make([]int, len(original[i]))
		for j := range original[i] {
			imageCopy[i][j] = original[i][j]
		}
	}
	return imageCopy
}

// generateFilledImage creates a rows×cols image filled with the specified color.
// This is useful for testing fill operations that cover the entire image.
func generateFilledImage(rows, cols, color int) [][]int {
	image := make([][]int, rows)
	for i := range image {
		image[i] = make([]int, cols)
		for j := range image[i] {
			image[i][j] = color
		}
	}
	return image
}

// generateMixedImage creates a rows×cols image with a mix of colors that better
// resembles a real-world scenario with distinct regions.
func generateMixedImage(rows, cols int) [][]int {
	image := make([][]int, rows)
	for i := range image {
		image[i] = make([]int, cols)
		for j := range image[i] {
			// Create a simple pattern with different colors
			if (i/4+j/4)%3 == 0 {
				image[i][j] = 1
			} else if (i/4+j/4)%3 == 1 {
				image[i][j] = 2
			} else {
				image[i][j] = 3
			}
		}
	}

	// Create a center region with a single color for flood filling
	centerColor := 4
	centerSize := 3
	centerRow := rows / 2
	centerCol := cols / 2

	for i := centerRow - centerSize; i <= centerRow+centerSize; i++ {
		for j := centerCol - centerSize; j <= centerCol+centerSize; j++ {
			if i >= 0 && i < rows && j >= 0 && j < cols {
				image[i][j] = centerColor
			}
		}
	}

	return image
}
