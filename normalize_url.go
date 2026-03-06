package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", nil
	}
	parsed.Scheme = ""
	final := strings.TrimPrefix(parsed.String(), "//")
	return final, nil
}
