package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: countWordsAndImages -url-\n")
		os.Exit(1)
	}
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("words: %d\nimages: %d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	stack := make([]*html.Node, 0)
	stack = append(stack, n)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop
		switch node.Type {
		case html.TextNode:
			words += countWords(node.Data)
		case html.ElementNode:
			if node.Data == "img" {
				images++
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			stack = append(stack, c)
		}
	}
	return
}

func countWords(s string) (words int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	return
}
