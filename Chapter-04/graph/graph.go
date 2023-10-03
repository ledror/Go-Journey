package graph

import "fmt"

type graphType map[string]map[string]bool

func NewGraph() graphType {
	return make(graphType)
}

func (g graphType) AddEdge(from, to string) {
	edges := g[from]
	if edges == nil {
		edges = make(map[string]bool)
		g[from] = edges
	}
	edges[to] = true
}

func (g graphType) HasEdge(from, to string) bool {
	return g[from][to]
}

func (g graphType) Print() {
	for src, edges := range g {
		fmt.Printf("%s -> { ", src)
		for dst, ok := range edges {
			if ok {
				fmt.Printf("%s, ", dst)
			}
		}
		fmt.Printf("}\n")
	}
}
