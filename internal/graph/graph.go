package graph

import (
	"errors"
	"fmt"
	"sort"
)

type Graph struct {
	nodes map[string]struct{}
	edges map[string]map[string]struct{}
}

func New() *Graph {
	return &Graph{
		nodes: map[string]struct{}{},
		edges: map[string]map[string]struct{}{},
	}
}

func (g *Graph) AddNode(id string) {
	if _, ok := g.nodes[id]; ok {
		return
	}
	g.nodes[id] = struct{}{}
}

func (g *Graph) HasNode(id string) bool {
	_, ok := g.nodes[id]
	return ok
}

func (g *Graph) Nodes() []string {
	out := make([]string, 0, len(g.nodes))
	for n := range g.nodes {
		out = append(out, n)
	}
	sort.Strings(out)
	return out
}

func (g *Graph) Successors(n string) []string {
	out := make([]string, 0, len(g.edges[n]))
	for to := range g.edges[n] {
		out = append(out, to)
	}
	sort.Strings(out)
	return out
}

func (g *Graph) Predecessors(n string) []string {
	var out []string
	for from, tos := range g.edges {
		if _, ok := tos[n]; ok {
			out = append(out, from)
		}
	}
	sort.Strings(out)
	return out
}

func (g *Graph) reaches(from, target string) bool {
	seen := map[string]bool{}
	stack := []string{from}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if n == target {
			return true
		}
		if seen[n] {
			continue
		}
		seen[n] = true
		for to := range g.edges[n] {
			stack = append(stack, to)
		}
	}
	return false
}

func (g *Graph) AddEdge(from, to string) error {
	if !g.HasNode(from) {
		return fmt.Errorf("edge from unknown node %q", from)
	}
	if !g.HasNode(to) {
		return fmt.Errorf("edge to unknown node %q", to)
	}
	if from == to {
		return errors.New("self-loop not allowed")
	}
	if g.reaches(to, from) {
		return fmt.Errorf("edge %q->%q would create a cycle", from, to)
	}
	if g.edges[from] == nil {
		g.edges[from] = map[string]struct{}{}
	}
	g.edges[from][to] = struct{}{}
	return nil
}

func (g *Graph) TopoSort() ([]string, error) {
	indeg := map[string]int{}
	for n := range g.nodes {
		indeg[n] = 0
	}
	for from := range g.edges {
		for to := range g.edges[from] {
			indeg[to]++
		}
	}
	visited := map[string]bool{}
	order := make([]string, 0, len(g.nodes))
	for len(visited) < len(g.nodes) {
		var ready []string
		for n := range g.nodes {
			if !visited[n] && indeg[n] == 0 {
				ready = append(ready, n)
			}
		}
		if len(ready) == 0 {
			return nil, errors.New("graph has a cycle")
		}
		sort.Strings(ready)
		n := ready[0]
		visited[n] = true
		order = append(order, n)
		for _, to := range g.Successors(n) {
			indeg[to]--
		}
	}
	return order, nil
}
