package model

type SearchResult struct {
	Hits  int
	Start int
	Songs []interface{}
}
