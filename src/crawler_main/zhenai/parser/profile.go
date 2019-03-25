package parser

import (
	"crawler_main/engine"
	"crawler_main/model"
	"regexp"
)

var jiguan = regexp.MustCompile(`<div class="m-btn pink" [^>]*>籍贯:([^<]+)</div>`)

func ParseProfile(
	contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name
	profile.Jiguan = extractString(
		contents, jiguan)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(
	contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return "aaa"
	}
}
