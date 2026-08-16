package impact

import (
	"sort"

	"etl-lineage/internal/graph"
	"etl-lineage/internal/lineage"
)

func Impact(g *graph.Graph, changed string) ([]string, error) {
	down, err := lineage.Downstream(g, changed)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0, len(down))
	for k := range down {
		res = append(res, k)
	}
	sort.Strings(res)
	return res, nil
}
