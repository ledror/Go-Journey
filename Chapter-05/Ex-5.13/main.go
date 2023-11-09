package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"Go-Journey/Chapter-05/links"
)

func main() {
	urlsChan = make(chan urlsOk)
	breadthFirstSearch(os.Args[1])
	log.Print("Done!")
}

type urlsOk struct {
	urls []string
	err  error
}

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
	if strings.HasPrefix(url, os.Args[1]) {
		copyPage(url)
	}
	links, err := links.Extract(url)
	if err != nil {
		urlsChan <- urlsOk{nil, err}
		return
	}
	urlsChan <- urlsOk{links, nil}
}

func copyPage(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("getting %s: %s", url, resp.Status)
		return
	}
	file_name := "tmp/" + path.Base(url)
	file, err := os.Create(file_name)
	if err != nil {
		log.Printf("error creating %s: %s", file_name, err)
		return
	}
	_, err = file.ReadFrom(resp.Body)
	if err != nil {
		log.Printf("error writing to %s: %s", file_name, err)
		return
	}
	if err := file.Close(); err != nil {
		log.Printf("error closing %s: %s", file_name, err)
		return
	}
}
