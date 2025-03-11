package trees

// TODO: implement constructor
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (r *TreeNode) String() string {
	// TODO: implement
	return "🌳"
}

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
