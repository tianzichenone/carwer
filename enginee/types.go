package enginee

type Request struct {
	URL        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    [] interface{}
}

func NilParserFunc(b []byte) ParserResult {
	return ParserResult{}
}


