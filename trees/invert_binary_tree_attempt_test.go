package trees_test

import (
	"fmt"
	"grind75/trees"
	"testing"
)

// TODO: implement test cases...
func TestInvertBinaryTree(t *testing.T) {
	cases := []struct {
		tree     *trees.TreeNode
		expected *trees.TreeNode
	}{
		{},
		{},
		{},
	}
	impls := map[string]trees.InvertTreeFunc{
		"InvertTreeAttempt": trees.InvertTreeAttempt,
	}

	for name, impl := range impls {
		for _, testCase := range cases {
			t.Run(fmt.Sprintf("%s (tree=%v)", name, testCase.tree), func(t *testing.T) {
				actual := impl(testCase.tree)
				if actual.String() != testCase.expected.String() {
					t.Errorf("wanted %v, got %v", testCase.expected, actual)
				}
			})
		}
	}
}
