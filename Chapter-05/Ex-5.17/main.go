package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) (matching []*html.Node) {
	var visit func(n *html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode && slices.Contains(name, n.Data) {
			matching = append(matching, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(doc)
	return
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	images := ElementsByTagName(doc, "img")
	for _, img := range images {
		fmt.Printf("<img %s>\n", img.Attr)
	}
}
