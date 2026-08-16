package lineage

import (
	"testing"

	"etl-lineage/internal/graph"
)

func build() *graph.Graph {
	g := graph.New()
	g.AddNode("a")
	g.AddNode("b")
	g.AddNode("c")
	g.AddEdge("a", "b")
	g.AddEdge("b", "c")
	return g
}

func TestUpstreamMissingNode(t *testing.T) {
	g := build()
	if _, err := Upstream(g, "ghost"); err == nil {
		t.Fatal("expected error for missing node")
	}
}

func TestUpstreamEmpty(t *testing.T) {
	g := build()
	up, err := Upstream(g, "a")
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if up == nil {
		t.Fatal("Upstream should return non-nil empty map for a node with no upstream")
	}
	if len(up) != 0 {
		t.Fatalf("want 0 upstream, got %d", len(up))
	}
}

func TestUpstreamEmptyMapNotNil(t *testing.T) {
	g := build()
	up, err := Upstream(g, "a")
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if up == nil {
		t.Fatal("root node with no predecessors must return a non-nil empty map, not nil")
	}
	if len(up) != 0 {
		t.Fatalf("want empty map, got %v", up)
	}
}

func TestUpstreamTransitive(t *testing.T) {
	g := build()
	up, err := Upstream(g, "c")
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if len(up) != 2 || !up["a"] || !up["b"] {
		t.Fatalf("upstream of c should be {a,b}, got %v", up)
	}
}

func TestDownstreamTransitive(t *testing.T) {
	g := build()
	down, err := Downstream(g, "a")
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if len(down) != 2 || !down["b"] || !down["c"] {
		t.Fatalf("downstream of a should be {b,c}, got %v", down)
	}
}

func TestDownstreamMissingNode(t *testing.T) {
	g := build()
	if _, err := Downstream(g, "ghost"); err == nil {
		t.Fatal("expected error for missing node")
	}
}
