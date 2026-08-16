package parse

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"etl-lineage/internal/graph"
)

func ParseSpec(r io.Reader) (*graph.Graph, error) {
	g := graph.New()
	sc := bufio.NewScanner(r)
	lineNo := 0
	for sc.Scan() {
		lineNo++
		raw := strings.TrimSpace(sc.Text())
		if raw == "" || strings.HasPrefix(raw, "#") {
			continue
		}
		idx := strings.Index(raw, "<-")
		if idx < 0 {
			return nil, fmt.Errorf("line %d: missing '<-' separator", lineNo)
		}
		target := strings.TrimSpace(raw[:idx])
		srcPart := strings.TrimSpace(raw[idx+2:])
		if target == "" {
			return nil, fmt.Errorf("line %d: empty target", lineNo)
		}
		g.AddNode(target)
		if srcPart == "" {
			continue
		}
		for _, dep := range strings.Split(srcPart, ",") {
			dep = strings.TrimSpace(dep)
			if dep == "" {
				continue
			}
			g.AddNode(dep)
			if err := g.AddEdge(dep, target); err != nil {
				return nil, fmt.Errorf("line %d: %w", lineNo, err)
			}
		}
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	return g, nil
}
