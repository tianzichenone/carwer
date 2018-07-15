package enginee

type Request struct{
	URL string
	parserFunc  func () RequestResult
}

type RequestResult struct {
	requests []Request
	items [] interface{}
}

