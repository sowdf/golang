package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"model"
	"qupu/engine"
	"qupu/fetcher"
	"regexp"
	"strconv"
)

var compositionRe = regexp.MustCompile(`<span>作曲：</span><a [^>]*>([^<]+)</a>`)
var idRe = regexp.MustCompile(`http://www.qupu123.com/[^p]*/p([\d]+).html`)
var fromRe = regexp.MustCompile(`<span>来源：</span>([^<]+)<span>`)
var lyricsRe = regexp.MustCompile(`<span>作词：</span><a [^>]*>([^<]+)</a>`)
var dateRe = regexp.MustCompile(`<span>日期：</span>([\d]+-[\d]+-[\d]+)`)
var uploaderRe = regexp.MustCompile(`<span>上传：</span>([^\n]+)\n`)
var imageListRe = regexp.MustCompile(`href="(/Public/Uploads/[^"]+)"`)

func ParseSong(contents []byte, url string) engine.ParseResult {

	newReader := bytes.NewReader(contents)
	document, err := goquery.NewDocumentFromReader(newReader)

	if err != nil {
		log.Printf("song页面选择器初始化失败：%s\n", err)
	}

	result := engine.ParseResult{}
	info, err := document.Find(".content_head").Html()
	if err != nil {
		log.Printf("info get Error %s", err)
	}

	song := model.Spectrum{}

	if err != nil {
		log.Printf("Views To Int Eroor %d\n", err)
	}
	song.Type = "qupu123"
	song.Composition = extractString(info, compositionRe)
	song.Lyrics = extractString(info, lyricsRe)
	song.Title = document.Find("h1").Text()
	song.Id = extractString(url, idRe)
	song.Views = getViews(song.Id)
	song.FromUrl = url
	song.From = extractString(info, fromRe)
	song.UploadDate = extractString(info, dateRe)
	song.Uploader = extractString(info, uploaderRe)
	song.ClickNum = 0

	imageListHtml, err := document.Find(".imageList").Html()

	if err != nil {
		log.Printf("imageList select Error %s", err)
	}

	allMatch := imageListRe.FindAllStringSubmatch(imageListHtml, -1)

	for _, m := range allMatch {
		song.ImageUrls = append(song.ImageUrls, "http://www.qupu123.com"+m[1])
	}

	result.Item = song

	return result
}

func getViews(id string) int {
	contents, err := fetcher.Fetch("http://www.qupu123.com/Opern-cnum-id-" + id + ".html")
	if err != nil {
		return 0
	}
	num, err := strconv.Atoi(string(contents))
	if err != nil {
		return 0
	}
	return num
}

func extractString(contents string, re *regexp.Regexp) string {
	match := re.FindStringSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
