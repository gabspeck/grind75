package trees

type InvertTreeFunc func(root *TreeNode) *TreeNode

func InvertTreeAttempt(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	InvertTreeAttempt(root.Left)
	InvertTreeAttempt(root.Right)

	return root
}
