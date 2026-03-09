package main

import (
	"log"
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	heading := getHeadingFromHTML(html)

	fParagraph := getFirstParagraphFromHTML(html)

	parsed, err := url.Parse(pageURL)
	if err != nil {
		log.Println(err)
	}
	links, err := getURLsFromHTML(html, parsed)
	if err != nil {
		log.Println(err)
	}

	imgs, err := getImagesFromHTML(html, parsed)
	if err != nil {
		log.Println(err)
	}

	return PageData{
		URL:            pageURL,
		Heading:        heading,
		FirstParagraph: fParagraph,
		OutgoingLinks:  links,
		ImageURLs:      imgs,
	}

}
