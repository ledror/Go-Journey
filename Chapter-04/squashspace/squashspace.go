package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpaces(b []byte) []byte {
	i := 0
	for i < len(b) {
		r, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(r) {
			for j := i + size; j < len(b); {
				r, size := utf8.DecodeRune(b[j:])
				if unicode.IsSpace(r) {
					j += size
					continue
				} else {
					b[i] = ' '
					i++
					copy(b[i:], b[j:])
					b = b[:len(b)-j+i]
					break
				}
			}
		} else {
			i += size
		}
	}
	return b
}

func main() {
	b := []byte("this has  many    \t  spaces")
	b_copy := []byte("this has many spaces")
	fmt.Printf("%s\ncap:%v\nlen:%v\n", b, cap(b), len(b))
	b = squashSpaces(b)
	fmt.Printf("%s\ncap:%v\nlen:%v\n", b, cap(b), len(b))
	fmt.Printf("%s\ncap:%v\nlen:%v\n", b, cap(b_copy), len(b_copy))
}
