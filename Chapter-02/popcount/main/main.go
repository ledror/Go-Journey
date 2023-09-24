package main

import (
	"fmt"
	"popcount"
)

func main() {
	x := uint64(0xffff)
	fmt.Printf("%d\n", popcount.PopCount(x))
	fmt.Printf("%d\n", popcount.PopCountLoop(x))
	fmt.Printf("%d\n", popcount.PopCountShifting(x))
	fmt.Printf("%d\n", popcount.PopCountFunny(x))
}
