package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, file_map := range counts {
		n := 0
		for _, i := range file_map {
			n += i
		}
		if n > 1 {
			fmt.Printf("Total=%d: \"%s\"\t", n, line)
			var sep string
			for file_name, i := range file_map {
				fmt.Printf("%s%s(%d)", sep, file_name, i)
				sep = ", "
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	file_name := f.Name()
	for input.Scan() {
		_, ok := counts[input.Text()]
		if !ok {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][file_name]++
	}
}
