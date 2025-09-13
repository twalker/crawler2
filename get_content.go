package main

import (
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
