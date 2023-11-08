package main

import (
	"fmt"
	"log"
	"os"

	"Go-Journey/Chapter-05/links"
)

func main() {
	breadthFirstSearch(crawl, os.Args[1:])
}

func breadthFirstSearch(expand func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, expand(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	links, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return links
}
