package engine

type Request struct {
	Url        string
	ParserFunc func(content []byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Item     interface{}
}
