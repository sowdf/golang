package parser

import (
	"crawker/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	//<a href="http://www.zhenai.com/zhenghun/xinzhu" data-v-5e16505f>新竹</a>
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	//[][]byte  这里[]byte 表示的是字符串
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: ParseCity})
	}
	return result
}
