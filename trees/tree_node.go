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
