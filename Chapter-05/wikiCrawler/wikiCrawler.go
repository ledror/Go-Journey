package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"

	"Go-Journey/Chapter-05/links"
)

const (
	base = "https://he.wikipedia.org/wiki/"
	src  = "https://he.wikipedia.org/wiki/כריסטיאנו_רונאלדו"
	dst  = "https://he.wikipedia.org/wiki/אטום"
)

func main() {
	urlsChan = make(chan urlsOk)
	startup = time.Now()
	breadthFirstSearch(src)
	log.Print("Done!")
}

type urlsOk struct {
	urls []string
	err  error
}

var (
	startup time.Time
	mu      sync.Mutex
	counter int
	depth   int
)

var UNRELATED_DOMAIN = errors.New("UNRELATED DOMAIN")

var urlsChan chan urlsOk

func breadthFirstSearch(root string) {
	worklist := []string{root}
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
				time.Sleep(10 * time.Millisecond) // 100 requests per second, too much will kick you
			}
		}
		depth++
		for i := 1; i <= n; i++ {
			resp := <-urlsChan
			if resp.err != nil {
				if resp.err != UNRELATED_DOMAIN {
					log.Print(resp.err)
				}
			} else {
				worklist = append(worklist, resp.urls...)
			}
		}
	}
}

func crawl(address string) {
	if !strings.HasPrefix(address, base) {
		urlsChan <- urlsOk{nil, UNRELATED_DOMAIN}
		return
	}
	mu.Lock()
	counter++
	mu.Unlock()
	decoded, err := url.QueryUnescape(address)
	if err != nil {
		urlsChan <- urlsOk{nil, err}
		return
	}
	fmt.Println(decoded)
	if decoded == dst {
		log.Fatalf(
			"Done!\nSearched \"only\" %d wiki pages and finished in just %v\nDepth: %d\n",
			counter,
			time.Now().Sub(startup),
			depth,
		)
	}
	links, err := links.Extract(decoded)
	if err != nil {
		urlsChan <- urlsOk{nil, err}
		return
	}
	urlsChan <- urlsOk{links, nil}
}
