package linkedlist

import "testing"

func TestHasCycle_NoCycle(t *testing.T) {
	head := FromSlice([]int{1, 2, 3, 4})
	if HasCycle(head) {
		t.Fatal("expected no cycle in a plain acyclic list")
	}
}

func TestHasCycle_EmptyList(t *testing.T) {
	if HasCycle(nil) {
		t.Fatal("expected no cycle in an empty list")
	}
}

func TestHasCycle_SingleNodeSelfLoop(t *testing.T) {
	n := &Node{Val: 1}
	n.Next = n
	if !HasCycle(n) {
		t.Fatal("expected self-loop to be detected as a cycle")
	}
}

func TestHasCycle_CycleInTail(t *testing.T) {
	head := FromSlice([]int{1, 2, 3, 4, 5})

	// Manually wire the tail back to the 3rd node to form a cycle.
	tail := head
	var third *Node
	for i, n := 0, head; n != nil; i, n = i+1, n.Next {
		if i == 2 {
			third = n
		}
		tail = n
	}
	tail.Next = third

	if !HasCycle(head) {
		t.Fatal("expected cycle to be detected")
	}
}
