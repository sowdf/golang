package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"model"
	"qupu/engine"
)

func ParseSong(contents []byte) engine.ParseResult {

	newReader := bytes.NewReader(contents)
	document, err := goquery.NewDocumentFromReader(newReader)

	if err != nil {
		log.Printf("song页面选择器初始化失败：%s\n", err)
	}

	result := engine.ParseResult{}

	song := model.Spectrum{
		Id:      "",
		Title:   document.Find(".content_head h1").Text(),
		From:    "",
		FromUrl: "",
	}

	result.Item = song

	return engine.ParseResult{}
}
