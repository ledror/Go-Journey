package main

import (
	"fmt"
)

func max(x int, vals ...int) int {
	curr := x
	for _, val := range vals {
		if curr < val {
			curr = val
		}
	}
	return curr
}

func min(x int, vals ...int) int {
	curr := x
	for _, val := range vals {
		if val < curr {
			curr = val
		}
	}
	return curr
}

func main() {
	fmt.Println(max(5, 10, 123, 3, -512, 12))
	fmt.Println(min(100, -1, 51, 123, -512))
}
