package graph

import (
	"testing"
)

func TestAddEdgeUnknownNode(t *testing.T) {
	g := New()
	if err := g.AddEdge("x", "y"); err == nil {
		t.Fatal("expected error for edge between unknown nodes")
	}
}

func TestAddEdgeCycle(t *testing.T) {
	g := New()
	g.AddNode("a")
	g.AddNode("b")
	if err := g.AddEdge("a", "b"); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if err := g.AddEdge("b", "a"); err == nil {
		t.Fatal("expected error creating cycle")
	}
}

func TestTopoSortOrder(t *testing.T) {
	g := New()
	g.AddNode("target")
	g.AddNode("dep1")
	g.AddNode("dep2")
	g.AddEdge("dep1", "target")
	g.AddEdge("dep2", "target")
	order, err := g.TopoSort()
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if len(order) != 3 {
		t.Fatalf("order len=%d want 3", len(order))
	}
	pos := map[string]int{}
	for i, n := range order {
		pos[n] = i
	}
	if pos["dep1"] > pos["target"] || pos["dep2"] > pos["target"] {
		t.Fatalf("dependencies must precede target: %v", order)
	}
}

func TestTopoSortCycle(t *testing.T) {
	g := New()
	g.AddNode("a")
	g.AddNode("b")
	g.AddEdge("a", "b")
	g.AddEdge("b", "a") // would be cycle, but AddEdge rejects; force via internal? Use reaches check: add a->b then b->a rejected.
	// build a real cycle through a third node
	g.AddNode("c")
	g.AddEdge("b", "c")
	if err := g.AddEdge("c", "a"); err == nil {
		t.Fatal("expected cycle detection")
	}
}
