package main

import (
	"fmt"

	"github.com/sourcegraph/sourcegraph/lib/errors"
	"github.com/sourcegraph/tf-dag/dag"
)

func main() {
	g := buildGraph()
	if err := g.Walk(func(v dag.Vertex) error {
		fmt.Printf("visiting %d\n", v)
		return nil
	}); err != nil {
		fmt.Printf("error walking dag: %s", err.Error())
	}

	g = buildGraph()
	if err := g.Walk(func(v dag.Vertex) error {
		fmt.Printf("visiting %d\n", v)
		return errors.Newf("error walking: %d", v)
	}); err != nil {
		fmt.Printf("error walking dag: %s", err.Error())
	}
}

// buildGraph builds a simple graph where
// some vertices could be run in parallel
func buildGraph() dag.AcyclicGraph {
	var g dag.AcyclicGraph
	g.Add(1)
	g.Add(2)
	g.Add(3)
	g.Add(4)
	g.Connect(dag.BasicEdge(1, 2))
	g.Connect(dag.BasicEdge(1, 3))
	g.Connect(dag.BasicEdge(1, 4))

	return g
}
