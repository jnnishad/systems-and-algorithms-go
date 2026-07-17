package trees

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	//       3
	//      / \
	//     9  20
	//        / \
	//       15  7
	root := &Node{
		Val:  3,
		Left: &Node{Val: 9},
		Right: &Node{
			Val:   20,
			Left:  &Node{Val: 15},
			Right: &Node{Val: 7},
		},
	}

	got := LevelOrder(root)
	want := [][]int{{3}, {9, 20}, {15, 7}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("LevelOrder() = %v, want %v", got, want)
	}
}

func TestLevelOrder_NilTree(t *testing.T) {
	if got := LevelOrder(nil); got != nil {
		t.Fatalf("expected nil for empty tree, got %v", got)
	}
}

func TestLevelOrder_SingleNode(t *testing.T) {
	got := LevelOrder(&Node{Val: 42})
	want := [][]int{{42}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("LevelOrder(single) = %v, want %v", got, want)
	}
}
