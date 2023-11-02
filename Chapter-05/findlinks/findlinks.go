package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	main1()
}

func main1() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func main2() {
	freqs, err := elementNameFreqs(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	names := make([]string, 0)
	for name := range freqs {
		names = append(names, name)
	}
	sort.SliceStable(names, func(i, j int) bool {
		return freqs[names[i]] > freqs[names[j]]
	})
	for _, name := range names {
		fmt.Printf("%d %s\n", freqs[name], name)
	}
}

func main3() {
	if err := printText(os.Stdin, os.Stdout); err != nil {
		fmt.Printf("error in printText: %v", err)
		os.Exit(1)
	}
	return
}

type strKeyVal struct {
	key, val string
}

func visit(links []strKeyVal, n *html.Node) []strKeyVal {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, strKeyVal{a.Key, a.Val})
			}
		}
	} else if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, strKeyVal{a.Key, a.Val})
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}

func elementNameFreqs(r io.Reader) (map[string]int, error) {
	freqs := make(map[string]int)
	tokenizer := html.NewTokenizer(r)
	for {
		t := tokenizer.Next()
		if t == html.ErrorToken {
			break
		}
		name, _ := tokenizer.TagName()
		if len(name) > 0 {
			freqs[string(name)]++
		}
	}
	if err := tokenizer.Err(); err != io.EOF {
		return freqs, err
	}
	return freqs, nil
}

func printText(r io.Reader, w io.Writer) error {
	tokenizer := html.NewTokenizer(r)
	stack := make([]string, 0)
Tokenizing:
	for {
		switch tokenizer.Next() {
		case html.ErrorToken:
			break Tokenizing
		case html.StartTagToken:
			name, _ := tokenizer.TagName()
			stack = append(stack, string(name))
		case html.EndTagToken:
			stack = stack[:len(stack)-1]
		case html.TextToken:
			if len(stack) > 0 && (stack[len(stack)-1] == "script" || stack[len(stack)-1] == "style") {
				continue
			}
			txt := tokenizer.Text()
			if len(strings.TrimSpace(string(txt))) == 0 {
				continue
			}
			_, err := w.Write(txt)
			if err != nil {
				return err
			}
			if txt[len(txt)-1] != '\n' {
				w.Write([]byte(string("\n")))
			}
		}
	}
	if err := tokenizer.Err(); err != io.EOF {
		return err
	}
	return nil
}
