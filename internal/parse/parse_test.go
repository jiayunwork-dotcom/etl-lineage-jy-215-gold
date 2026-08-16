package parse

import (
	"strings"
	"testing"
)

func TestParseSpecOK(t *testing.T) {
	in := "a <- b, c\nb <- c\n"
	g, err := ParseSpec(strings.NewReader(in))
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if !g.HasNode("a") || !g.HasNode("b") || !g.HasNode("c") {
		t.Fatalf("nodes missing: %v", g.Nodes())
	}
	// a depends on b and c: edges c->a, b->a; b depends on c: edge c->b
	if len(g.Successors("c")) != 2 {
		t.Fatalf("c should have 2 outgoing edges, got %v", g.Successors("c"))
	}
}

func TestParseSpecMissingSeparator(t *testing.T) {
	if _, err := ParseSpec(strings.NewReader("a b c\n")); err == nil {
		t.Fatal("expected error for missing '<-' separator")
	}
}

func TestParseSpecEmptyTarget(t *testing.T) {
	if _, err := ParseSpec(strings.NewReader(" <- a\n")); err == nil {
		t.Fatal("expected error for empty target")
	}
}

func TestParseSpecRejectsCycle(t *testing.T) {
	// a <- b  then b <- a would create a cycle via edges b->a and a->b
	in := "a <- b\nb <- a\n"
	if _, err := ParseSpec(strings.NewReader(in)); err == nil {
		t.Fatal("expected error when spec edges form a cycle")
	}
}
