package engine

import (
	"crawker/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	body, e := fetcher.Fetch(r.Url)
	if e != nil {
		log.Printf("Fetcher : error fetching url %s:%v")
		return ParseResult{}, e
	}
	return r.ParserFunc(body, r.Url), nil
}
