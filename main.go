package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"etl-lineage/internal/impact"
	"etl-lineage/internal/parse"
	"etl-lineage/internal/report"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "etl-lineage:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("etl-lineage", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	spec := fs.String("spec", "-", "lineage spec file ('-' for stdin)")
	query := fs.String("node", "", "node to compute downstream impact for")
	dot := fs.String("dot", "-", "write DOT graph ('-' for stdout)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	var r io.Reader = os.Stdin
	if *spec != "-" {
		f, err := os.Open(*spec)
		if err != nil {
			return fmt.Errorf("open spec: %w", err)
		}
		defer f.Close()
		r = f
	}
	g, err := parse.ParseSpec(r)
	if err != nil {
		return fmt.Errorf("parse spec: %w", err)
	}
	if *query != "" {
		imp, err := impact.Impact(g, *query)
		if err != nil {
			return fmt.Errorf("impact: %w", err)
		}
		for _, n := range imp {
			fmt.Println(n)
		}
		return nil
	}
	var w io.Writer = os.Stdout
	if *dot != "-" {
		of, err := os.Create(*dot)
		if err != nil {
			return fmt.Errorf("create dot: %w", err)
		}
		defer of.Close()
		w = of
	}
	if err := report.WriteDOT(w, g); err != nil {
		return fmt.Errorf("write dot: %w", err)
	}
	return nil
}
