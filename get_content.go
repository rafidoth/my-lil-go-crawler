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

func getFirstParagraphFromHTML(rawHTML string) string {
	doc, err := html.Parse(strings.NewReader(rawHTML))
	if err != nil {
		panic(err)
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.Main {
			for n1 := range n.Descendants() {
				if n1.Type == html.ElementNode && n1.DataAtom == atom.P {
					if n1.FirstChild != nil && n1.FirstChild.Type == html.TextNode {
						return n1.FirstChild.Data
					}
				}
			}
		}
	}
	return ""
}
