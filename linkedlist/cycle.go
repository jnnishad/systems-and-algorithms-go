package linkedlist

// HasCycle detects whether a linked list contains a cycle using
// Floyd's tortoise-and-hare algorithm: O(n) time, O(1) space — no
// auxiliary set of visited nodes required.
func HasCycle(head *Node) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
