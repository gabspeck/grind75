package trees

// Given the root of a binary tree, return the length of the diameter of the tree.
//
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
//
// The length of a path between two nodes is represented by the number of edges between them.
//
// Constraints:
//
// The number of nodes in the tree is in the range [1, 10^4].
//
// -100 <= Node.val <= 100
type DiameterOfBinaryTreeFunc func(*TreeNode) int

func getNodeHeight(root *TreeNode, treeDiameter *int) int {

	if root == nil {
		return 0
	}

	leftHeight := getNodeHeight(root.Left, treeDiameter)
	rightHeight := getNodeHeight(root.Right, treeDiameter)

	nodeDiameter := leftHeight + rightHeight

	if nodeDiameter > *treeDiameter {
		*treeDiameter = nodeDiameter
	}

	return max(leftHeight, rightHeight) + 1

}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	getNodeHeight(root, &diameter)

	return diameter
}
