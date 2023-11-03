package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: wordfreq -filenmae-\n")
		return
	}
	file, err := os.Open(strings.Join(os.Args[1:], " "))
	if err != nil {
		fmt.Printf("wordfreq: error opening %s", strings.Join(os.Args[1:], " "))
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	freqs := make(map[string]int)
	for scanner.Scan() {
		freqs[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("reading file: %v\n", err)
		return
	}
	fmt.Printf("word\tcount\n")
	for word, count := range freqs {
		fmt.Printf("%s\t%d\n", word, count)
	}
}
