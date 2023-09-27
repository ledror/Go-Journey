package main

import (
	"fmt"

	"ch4/reverse"
)

func main() {
	// a := [...]int{0, 1, 2, 3, 4, 5}
	// reverse.Reverse(a[:2])
	// reverse.Reverse(a[2:])
	// reverse.Reverse(a[:])
	// ints := []int{1, 2, 3, 4, 5}
	// ints = reverse.Rotate(ints, 2)
	// fmt.Println(ints)
	s := "12😿34asdf"
	reversed_s := string(reverse.ReverseUTF([]byte(string(s))))
	fmt.Println(reversed_s)
}
