package lineage

import (
	"errors"

	"etl-lineage/internal/graph"
)

func Upstream(g *graph.Graph, id string) (map[string]bool, error) {
	if !g.HasNode(id) {
		return nil, errors.New("node not found")
	}
	res := map[string]bool{}
	var dfs func(string)
	dfs = func(n string) {
		for _, from := range g.Predecessors(n) {
			if !res[from] {
				res[from] = true
				dfs(from)
			}
		}
	}
	dfs(id)
	delete(res, id)
	return res, nil
}

func Downstream(g *graph.Graph, id string) (map[string]bool, error) {
	if !g.HasNode(id) {
		return nil, errors.New("node not found")
	}
	res := map[string]bool{}
	var dfs func(string)
	dfs = func(n string) {
		for _, to := range g.Successors(n) {
			if !res[to] {
				res[to] = true
				dfs(to)
			}
		}
	}
	dfs(id)
	delete(res, id)
	return res, nil
}
