package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"qupu/engine"
	"regexp"
)

var songUrlRe = regexp.MustCompile(`<a href="([^"]+)" [^>]*>([^<])+</a>`)
var categoryLinkRe = regexp.MustCompile(`href="([^"]+)"`)

func ParseCategory(contents []byte) engine.ParseResult {
	newReader := bytes.NewReader(contents)
	document, err := goquery.NewDocumentFromReader(newReader)
	if err != nil {
		log.Printf("分类页面选择器初始化失败%s\n", err)
	}
	html, err := document.Find(".opern_list").Html()

	if err != nil {
		log.Printf("分类选择器初始化失败%s\n", err)
	}

	allMatch := songUrlRe.FindAllStringSubmatch(html, -1)
	result := engine.ParseResult{}
	for _, m := range allMatch {
		url := "http://www.qupu123.com" + m[1]
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(content []byte) engine.ParseResult {
				return ParseSong(contents, url)
			},
		})
	}

	html, err = document.Find(".pageHtml").Html()

	if err != nil {
		log.Printf("分类页面分页选择器查找失败%s\n", err)
	}

	categoryLinkMatch := categoryLinkRe.FindAllStringSubmatch(html, -1)

	for _, m := range categoryLinkMatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:        "http://www.qupu123.com" + m[1],
			ParserFunc: ParseCategory,
		})
	}

	return result

}
