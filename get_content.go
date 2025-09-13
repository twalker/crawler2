package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) (h1text string) {
	r := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return ""
	}
	h1 := doc.Find("h1").First()
	if h1.Length() == 0 {
		return ""
	}
	return strings.TrimSpace(h1.Text())
}

func getFirstParagraphFromHTML(html string) (paragraphText string) {
	r := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return ""
	}
	p := doc.Find("p").First()
	if p.Length() == 0 {
		return ""
	}
	return strings.TrimSpace(p.Text())
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	var urls []string
	r := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	doc.Find("a,img").Each(func(i int, s *goquery.Selection) {
		if a, ok := s.Attr("src"); ok {
			urls = append(urls, a)
		}
		if a, ok := s.Attr("href"); ok {
			urls = append(urls, a)
		}
	})

	for i, u := range urls {
		if !strings.HasPrefix(u, "http") {
			urls[i] = baseURL.String() + u
		}
	}

	return urls, nil
}
