package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

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
		if exists && s.Is("img") {
			processedURL, err := processUrl(requestBase, url)
			if err != nil {
				log.Fatal(err)
			}
			urls = append(urls, processedURL)
		}
	})

	err = downloadImages(urls)
	if err != nil {
		log.Fatal(err)
	}
}

// processUrl converts relative urls to absolute urls
func processUrl(base, path string) (string, error) {
	switch path[0] {
	case '/':
		return base + path, nil
	case '.':
		parse, err := url.Parse(base)
		if err != nil {
			return "", err
		}
		return parse.JoinPath(path).String(), nil
	default:
		return path, nil
	}
}

func getUrlBase(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return parsedUrl.Scheme + "://" + parsedUrl.Host, nil
}

func downloadImages(urls []string) error {
	err := os.Mkdir("images", 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}
	err = os.Chdir("images")
	if err != nil {
		return err
	}
	for _, url := range urls {
		base := path.Base(url)
		f, err := os.Create(base)
		if err != nil {
			log.Default().Println(err)
			continue
		}
		defer f.Close()
		resp, err := http.Get(url)
		if err != nil {
			log.Default().Println(err)
			continue
		}
		io.Copy(f, resp.Body)
	}
	return nil
}
