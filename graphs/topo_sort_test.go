package graphs

import "testing"

// indexOf returns the position of val in order, or -1.
func indexOf(order []string, val string) int {
	for i, v := range order {
		if v == val {
			return i
		}
	}
	return -1
}

// assertBefore fails the test if a does not appear before b in order.
func assertBefore(t *testing.T, order []string, a, b string) {
	t.Helper()
	ia, ib := indexOf(order, a), indexOf(order, b)
	if ia == -1 || ib == -1 {
		t.Fatalf("expected both %q and %q in order %v", a, b, order)
	}
	if ia >= ib {
		t.Fatalf("expected %q before %q, got order %v", a, b, order)
	}
}

func TestTopoSort_SimpleChain(t *testing.T) {
	graph := map[string][]string{
		"c": {"b"},
		"b": {"a"},
		"a": {},
	}
	order, err := TopoSort(graph)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertBefore(t, order, "a", "b")
	assertBefore(t, order, "b", "c")
}

// This mirrors a real GitOps dependency graph: a Helm release for a
// web service depends on its database and config being applied first,
// and the database depends on the namespace existing.
func TestTopoSort_GitOpsStyleDependencyGraph(t *testing.T) {
	graph := map[string][]string{
		"web-service": {"database", "config-map"},
		"database":    {"namespace"},
		"config-map":  {"namespace"},
		"namespace":   {},
	}

	order, err := TopoSort(graph)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assertBefore(t, order, "namespace", "database")
	assertBefore(t, order, "namespace", "config-map")
	assertBefore(t, order, "database", "web-service")
	assertBefore(t, order, "config-map", "web-service")

	if len(order) != 4 {
		t.Fatalf("expected all 4 nodes in the order, got %v", order)
	}
}

func TestTopoSort_DetectsDirectCycle(t *testing.T) {
	graph := map[string][]string{
		"a": {"b"},
		"b": {"a"},
	}
	_, err := TopoSort(graph)
	if err == nil {
		t.Fatal("expected an error for a 2-node cycle")
	}
}

func TestTopoSort_DetectsSelfCycle(t *testing.T) {
	graph := map[string][]string{
		"a": {"a"},
	}
	_, err := TopoSort(graph)
	if err == nil {
		t.Fatal("expected an error for a self-referential dependency")
	}
}

func TestTopoSort_DependencyNotListedAsKey(t *testing.T) {
	// "external-base" is a dependency but never appears as its own key —
	// this is the common case of depending on something outside the repo.
	graph := map[string][]string{
		"app": {"external-base"},
	}
	order, err := TopoSort(graph)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertBefore(t, order, "external-base", "app")
}

func TestTopoSort_EmptyGraph(t *testing.T) {
	order, err := TopoSort(map[string][]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(order) != 0 {
		t.Fatalf("expected empty order, got %v", order)
	}
}
