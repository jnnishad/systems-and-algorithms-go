package trees

// LevelOrder returns the tree's node values grouped by depth
// (breadth-first / BFS traversal), root level first.
func LevelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			n := queue[0]
			queue = queue[1:]
			level = append(level, n.Val)

			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}

		result = append(result, level)
	}

	return result
}
