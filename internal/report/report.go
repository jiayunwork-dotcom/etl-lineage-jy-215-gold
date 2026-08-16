package report

import (
	"bufio"
	"fmt"
	"io"

	"etl-lineage/internal/graph"
)

func WriteDOT(w io.Writer, g *graph.Graph) (err error) {
	bw := bufio.NewWriter(w)
	defer func() {
		if ferr := bw.Flush(); ferr != nil && err == nil {
			err = ferr
		}
	}()
	fmt.Fprintln(bw, "digraph lineage {")
	for _, n := range g.Nodes() {
		fmt.Fprintf(bw, "  %q;\n", n)
	}
	for _, from := range g.Nodes() {
		for _, to := range g.Successors(from) {
			fmt.Fprintf(bw, "  %q -> %q;\n", from, to)
		}
	}
	fmt.Fprintln(bw, "}")
	return nil
}
