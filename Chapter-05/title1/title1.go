package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/html" && !strings.HasPrefix(contentType, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, contentType)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var found bool
	var visit func(n *html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
			found = true
		}
		for c := n.FirstChild; c != nil && found == false; c = c.NextSibling {
			visit(c)
		}
	}
	visit(doc)
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: title1 <url>\n")
		return
	}
	err := title(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
}
