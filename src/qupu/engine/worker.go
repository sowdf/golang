package engine

import (
	"log"
	"qupu/fetcher"
)

func worker(r Request) (ParseResult, error) {
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher : err fetchng url %s\n", err)
		return ParseResult{}, err
	}

	return r.ParserFunc(contents, r.Url), nil
}
