package main

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getHeadingFromHTML(rawHTML string) string {
	doc, err := html.Parse(strings.NewReader(rawHTML))
	if err != nil {
		panic(err)
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.H1 {
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				return n.FirstChild.Data
			}
		}
	}
	return ""
}
