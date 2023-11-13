package main

import (
	"fmt"
)

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics when x==0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
