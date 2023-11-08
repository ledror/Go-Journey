package main

import (
	"fmt"
	"os"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

type parentNode struct {
	s      string
	parent *parentNode
}

func parentCallBack(n *parentNode) string {
	var path []string
	var tracePath func(node *parentNode)
	tracePath = func(node *parentNode) {
		if node == nil {
			return
		}
		tracePath(node.parent)
		path = append(path, node.s)
	}
	tracePath(n)
	return strings.Join(path, " -> ")
}

func topoSort(m map[string][]string) (order []string, err error) {
	marked := make(map[string]bool)
	var visit func(items []string, parent *parentNode)
	visit = func(items []string, parent *parentNode) {
		if err != nil {
			return
		}
		for _, v := range items {
			isMarked, isSeen := marked[v]
			if isMarked {
				return
			}
			if isSeen {
				err = fmt.Errorf("cycle detected: %s", parentCallBack(&parentNode{v, parent}))
				return
			} else {
				marked[v] = false
				tmp := &parentNode{v, parent}
				visit(m[v], tmp)
				marked[v] = true
				order = append(order, v)
			}
		}
	}
	for k := range m {
		visit([]string{k}, nil)
	}
	return order, err
}
