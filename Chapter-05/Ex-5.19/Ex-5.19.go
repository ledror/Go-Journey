package main

import (
	"fmt"
)

func main() {
	fmt.Println(f(5))
}

func f(x int) (val int) {
	defer func() {
		if p := recover(); p != nil {
			val = p.(int)
		}
	}()
	panic(x * x)
}
