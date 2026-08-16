package impact

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

func TestImpactSorted(t *testing.T) {
	g := build()
	got, err := Impact(g, "a")
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	want := []string{"b", "c"}
	if len(got) != 2 || got[0] != "b" || got[1] != "c" {
		t.Fatalf("impact of a = %v, want %v", got, want)
	}
}

func TestImpactMissingNode(t *testing.T) {
	g := build()
	if _, err := Impact(g, "ghost"); err == nil {
		t.Fatal("expected error for missing node")
	}
}

func TestImpactIndependentResults(t *testing.T) {
	g := build()
	first, err := Impact(g, "a")
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if len(first) != 2 || first[0] != "b" || first[1] != "c" {
		t.Fatalf("first impact = %v, want [b c]", first)
	}
	// hold onto first; a second Impact must not rewrite its backing store
	firstCopy := append([]string(nil), first...)
	second, err := Impact(g, "b")
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if len(second) != 1 || second[0] != "c" {
		t.Fatalf("second impact = %v, want [c]", second)
	}
	if first[0] != firstCopy[0] || first[1] != firstCopy[1] || len(first) != len(firstCopy) {
		t.Fatalf("first result mutated after second Impact: was %v, now %v", firstCopy, first)
	}
}
