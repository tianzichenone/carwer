package enginee

import "carwer/model"

type Request struct {
	URL        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    []model.Item
}

func NilParserFunc(b []byte) ParserResult {
	return ParserResult{}
}


