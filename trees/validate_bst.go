package trees

import "math"

// IsValidBST reports whether the tree rooted at root satisfies the
// binary-search-tree invariant: every node in a left subtree is less
// than its ancestor, every node in a right subtree is greater.
//
// The common bug here is checking only the immediate parent-child
// relationship; that misses cases like a right-left grandchild that's
// smaller than the root but larger than its immediate parent. This
// carries a valid (min, max) range down through the whole subtree
// instead, so every node is checked against every ancestor that
// constrains it, not just its direct parent.
func IsValidBST(root *Node) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(n *Node, min, max int) bool {
	if n == nil {
		return true
	}
	if n.Val <= min || n.Val >= max {
		return false
	}
	return validate(n.Left, min, n.Val) && validate(n.Right, n.Val, max)
}
