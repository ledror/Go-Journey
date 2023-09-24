package main

import (
	"fmt"
	"sort"
)

func main() {
	var s1, s2 string
	fmt.Print("Enter first string: ")
	fmt.Scan(&s1)
	fmt.Print("Enter second string: ")
	fmt.Scan(&s2)
	if isAnagram(s1, s2) {
		fmt.Printf("%s and %s are anagrams\n", s1, s2)
	} else {
		fmt.Printf("%s and %s aren't anagrams\n", s1, s2)
	}
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	r1, r2 := []rune(s1), []rune(s2)
	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})
	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})
	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}
