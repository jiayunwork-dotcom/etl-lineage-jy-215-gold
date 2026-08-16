package report

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"etl-lineage/internal/graph"
)

type failWriter struct{ fail bool }

func (f failWriter) Write(p []byte) (int, error) {
	if f.fail {
		return 0, errors.New("write fail")
	}
	return len(p), nil
}

func TestWriteDOTOK(t *testing.T) {
	g := graph.New()
	g.AddNode("a")
	g.AddNode("b")
	g.AddEdge("a", "b")
	var buf bytes.Buffer
	if err := WriteDOT(&buf, g); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if !strings.Contains(buf.String(), "digraph lineage") {
		t.Fatalf("missing content: %q", buf.String())
	}
}

func TestWriteDOTFlushError(t *testing.T) {
	g := graph.New()
	g.AddNode("a")
	if err := WriteDOT(failWriter{true}, g); err == nil {
		t.Fatal("expected flush error to propagate")
	}
}
