package engine

import (
	"model"
)

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url    string
	Parser Parser
}

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type NilParser struct {
}

func (*NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (*NilParser) Serialize() (name string, args interface{}) {
	return "NirParser", nil
}

type ParseResult struct {
	Requests []Request
	Item     model.Spectrum
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

//工厂行数

func NewFuncParser(parser ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: parser,
		name:   name,
	}
}
