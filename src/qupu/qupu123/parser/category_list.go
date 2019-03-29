package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"qupu/engine"
	"regexp"
)

var categoryUrlRe = regexp.MustCompile(`<a href="([^"]+)">([^<]+)</a>`)

func ParseCategoryList(contents []byte, _ string) engine.ParseResult {
	newReader := bytes.NewReader(contents)
	document, err := goquery.NewDocumentFromReader(newReader)
	if err != nil {
		log.Printf("选择器文档初始化失败：%s", err)
	}

	html, err := document.Find(".navigationChild").Html()

	if err != nil {
		log.Printf("选择元素失败：%s", err)
	}

	allMatch := categoryUrlRe.FindAllStringSubmatch(html, -1)
	result := engine.ParseResult{}

	for _, match := range allMatch {
		url := "http://www.qupu123.com" + match[1]
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: ParseCategory,
		})
	}

	return result
}
