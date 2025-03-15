package trees

// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
type LowestCommonAncestorFunc func(root, p, q *TreeNode) *TreeNode

func LowestCommonAncestorAttempt(root, p, q *TreeNode) *TreeNode {

	n := root

	for n != nil {
		if n.Val < p.Val && n.Val < q.Val {
			n = n.Right
		} else if n.Val > p.Val && n.Val > q.Val {
			n = n.Left
		} else {
			break
		}
	}

	return n

}
