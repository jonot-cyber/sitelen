package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://www.omniglot.com/conscripts/sitelenpona.htm")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		url, exists := s.Attr("src")
		if exists {
			fmt.Println(url)
		}
	})
}
