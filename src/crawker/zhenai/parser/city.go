package parser

import (
	"crawker/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
var cityRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	//[][]byte  这里[]byte 表示的是字符串
	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: func(c []byte) engine.ParseResult {
			return ParseProfile(c, name)
		}})
	}

	submatch := cityRe.FindAllSubmatch(contents, -1)
	for _, m := range submatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
