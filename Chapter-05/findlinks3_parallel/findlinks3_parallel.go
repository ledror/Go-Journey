package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"Go-Journey/Chapter-05/links"
)

func main() {
	urlsChan = make(chan urlsOk)
	breadthFirstSearch(os.Args[1:])
	log.Print("Done!")
}

type urlsOk struct {
	urls []string
	err  error
}

var urlsChan chan urlsOk

func breadthFirstSearch(worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		n := 0
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				n++
				go crawl(item)
			}
			time.Sleep(
				10 * time.Millisecond,
			) // websites dont like getting trillion requests from the same ip in a span of a nanosecond
		}
		for i := 1; i <= n; i++ {
			resp := <-urlsChan
			if resp.err != nil {
				log.Print(resp.err)
			} else {
				worklist = append(worklist, resp.urls...)
			}
		}
	}
}

func crawl(url string) {
	fmt.Println(url)
	links, err := links.Extract(url)
	if err != nil {
		urlsChan <- urlsOk{nil, err}
		return
	}
	urlsChan <- urlsOk{links, nil}
}
