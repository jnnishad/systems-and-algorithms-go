package trees

import "testing"

func TestIsValidBST_ValidTree(t *testing.T) {
	// 2
	// / \
	// 1   3
	root := &Node{Val: 2, Left: &Node{Val: 1}, Right: &Node{Val: 3}}
	if !IsValidBST(root) {
		t.Fatal("expected a valid BST to pass")
	}
}

func TestIsValidBST_NilTreeIsValid(t *testing.T) {
	if !IsValidBST(nil) {
		t.Fatal("expected an empty tree to be a valid BST")
	}
}

func TestIsValidBST_DirectParentChildViolation(t *testing.T) {
	// 5
	// / \
	// 1   4   <- 4 < 5 but is in the right subtree: invalid
	root := &Node{Val: 5, Left: &Node{Val: 1}, Right: &Node{Val: 4}}
	if IsValidBST(root) {
		t.Fatal("expected right child smaller than root to be invalid")
	}
}

// This is the classic trap: node 6 is less than its immediate parent
// 15 (6 < 15, looks fine on a naive "compare to direct parent only"
// check) but 6 is in node 10's *right* subtree, which means it must
// also be greater than 10 -- and it isn't. A parent-only check would
// wrongly accept this tree; the range-based check must reject it.
func TestIsValidBST_GrandchildViolatesAncestorNotParent(t *testing.T) {
	//      10
	//     /  \
	//    5    15
	//        /  \
	//       6    20
	root := &Node{
		Val:  10,
		Left: &Node{Val: 5},
		Right: &Node{
			Val:   15,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 20},
		},
	}
	if IsValidBST(root) {
		t.Fatal("expected grandchild that violates an ancestor's bound to be invalid")
	}
}

func TestIsValidBST_DuplicateValuesAreInvalid(t *testing.T) {
	root := &Node{Val: 2, Left: &Node{Val: 2}}
	if IsValidBST(root) {
		t.Fatal("expected duplicate value to be invalid (strict BST)")
	}
}
