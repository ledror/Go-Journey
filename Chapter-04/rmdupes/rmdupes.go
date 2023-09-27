package main

import "fmt"

func removeAdjacentDuplicates(strings []string) []string {
	i := 0
	for _, s := range strings {
		if strings[i] == s {
			continue
		}
		i++
		strings[i] = s
	}
	return strings[:i+1]
}

func main() {
	strings := []string{"1", "1", "1", "asdf", "asdf", "abc", "def", "ghi", "ghi"}
	strings = removeAdjacentDuplicates(strings)
	fmt.Println(strings)
}
