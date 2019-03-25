package engine

import (
	"crawker/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	//遍历每个一种子
	for _, r := range seeds {
		requests = append(requests, r)
	}
	//如果种子 > 0 就要继续执行
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, e := worker(r)
		if e != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

	}
}

func worker(r Request) (ParseResult, error) {
	body, e := fetcher.Fetch(r.Url)
	if e != nil {
		log.Printf("Fetcher : error fetching url %s:%v")
		return ParseResult{}, e
	}
	return r.ParserFunc(body), nil
}
