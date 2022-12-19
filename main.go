package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var requestURL = "https://www.omniglot.com/conscripts/sitelenpona.htm"
	requestBase, err := getUrlBase(requestURL)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocument(requestURL)
	if err != nil {
		log.Fatal(err)
	}

	var urls []string
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		url, exists := s.Attr("src")
		if exists {
			urls = append(urls, processUrl(requestBase, url))
		}
	})

	for _, url := range urls {
		fmt.Println(url)
	}
}

// processUrl converts relative urls to absolute urls
func processUrl(base, url string) string {
	if strings.HasPrefix(url, "/") {
		// relative
		return base + url
	}
	return url
}

func getUrlBase(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return parsedUrl.Scheme + "://" + parsedUrl.Host, nil
}
