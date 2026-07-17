package linkedlist

// Reverse reverses a singly linked list in place and returns the new
// head. O(n) time, O(1) space — three-pointer iterative reversal.
func Reverse(head *Node) *Node {
	var prev *Node
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
