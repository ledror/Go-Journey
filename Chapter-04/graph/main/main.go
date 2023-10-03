package main

import (
	"ch4/graph"
)

func main() {
	g := graph.NewGraph()
	g.AddEdge("a", "b")
	g.AddEdge("b", "a")
	g.Print()
}
