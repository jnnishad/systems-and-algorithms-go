// Package graphs contains graph algorithms — the kind that show up
// directly in infra tooling: topological sort is how you'd sequence
// Terraform module applies or Helm chart dependencies, not just a
// whiteboard exercise.
package graphs

import "fmt"

// TopoSort returns a valid topological ordering of the nodes in graph,
// where graph[node] lists that node's dependencies (edges that must
// come before it). Returns an error if the graph contains a cycle,
// since no valid ordering exists in that case.
//
// Implemented as Kahn's algorithm (BFS via in-degree counting) rather
// than DFS-with-recursion-stack: O(V+E) time, and it naturally detects
// cycles by noticing not every node got processed, without needing a
// separate "currently visiting" color state.
func TopoSort(graph map[string][]string) ([]string, error) {
	inDegree := make(map[string]int)
	for node := range graph {
		if _, ok := inDegree[node]; !ok {
			inDegree[node] = 0
		}
		for _, dep := range graph[node] {
			inDegree[node]++
			if _, ok := inDegree[dep]; !ok {
				inDegree[dep] = 0
			}
		}
	}

	// Nodes with no unresolved dependencies can go first. Sort the
	// initial queue for deterministic output across runs.
	var queue []string
	for node, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, node)
		}
	}
	sortStrings(queue)

	// dependents[d] = nodes that depend on d, i.e. reverse edges — used
	// to know whose in-degree to decrement once d is resolved.
	dependents := make(map[string][]string)
	for node, deps := range graph {
		for _, dep := range deps {
			dependents[dep] = append(dependents[dep], node)
		}
	}

	var order []string
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)

		var freed []string
		for _, dependent := range dependents[node] {
			inDegree[dependent]--
			if inDegree[dependent] == 0 {
				freed = append(freed, dependent)
			}
		}
		sortStrings(freed)
		queue = append(queue, freed...)
	}

	if len(order) != len(inDegree) {
		return nil, fmt.Errorf("graph has a cycle: only resolved %d of %d nodes", len(order), len(inDegree))
	}

	return order, nil
}

// sortStrings is a tiny insertion sort — avoids importing "sort" for a
// package that otherwise has zero dependencies, and these slices are
// small (a node's immediate ready-set, not the whole graph).
func sortStrings(s []string) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}
