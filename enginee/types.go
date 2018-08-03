package enginee

import "carwer/model"

type ParserFunc func(contents []byte, url string) ParserResult

type Request struct {
	URL        string
	ParserFunc ParserFunc
}

type ParserResult struct {
	Requests []Request
	Items    []model.Item
}

func NilParserFunc(b []byte) ParserResult {
	return ParserResult{}
}


