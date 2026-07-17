// Package linkedlist contains singly-linked-list problems: the pointer
// manipulation classics that show up disproportionately often in
// interview loops relative to how often you'd actually reach for a
// linked list in production Go.
package linkedlist

type Node struct {
	Val  int
	Next *Node
}

// FromSlice builds a linked list from a slice, for convenient test setup.
func FromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// ToSlice flattens a linked list back into a slice, for convenient
// assertions in tests.
func ToSlice(head *Node) []int {
	var out []int
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}
