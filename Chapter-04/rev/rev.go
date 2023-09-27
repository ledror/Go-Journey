package reverse

import (
	"fmt"
	"unicode/utf8"
)

func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ReverseArray(s *[4]int) {
	for i, j := 0, len(s); i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// rotates n times to the left
func Rotate(ints []int, n int) []int {
	temp := make([]int, n) // var temp []int
	copy(temp, ints[:n])   // temp = append(temp, ints[:n]...)
	fmt.Println(temp)
	ints = ints[n:]
	ints = append(ints, temp...)
	return ints
}

func ReverseUTF(b []byte) []byte {
	out := make([]byte, len(b))
	temp := b
	for len(temp) > 0 {
		r, count := utf8.DecodeRune(temp)
		copy(out[len(temp)-count:len(temp)], []byte(string(r)))
		temp = temp[count:]
	}
	copy(b, out)
	return b
}
