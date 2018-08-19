package enginee

import (
	"carwer/model"
)

type ParserFunc func(contents []byte, url string) ParserResult

type Parser interface {
	Serialize() (name string, args interface{})
	Parse(contents []byte, url string) ParserResult
}

type Request struct {
	URL    string
	Parser Parser
}

type ParserResult struct {
	Requests []Request
	Items    []model.Item
}

type ParserFuncFactory struct {
	Name           string
	RealParserFunc ParserFunc
}

func (p *ParserFuncFactory) Parse(contents []byte, url string) ParserResult {
	return p.RealParserFunc(contents, url)
}

func (p *ParserFuncFactory) Serialize() (string, interface{}) {
	return p.Name, nil
}

func NewParserFuncFactory(name string, p ParserFunc) *ParserFuncFactory {
	return &ParserFuncFactory{
		Name:           name,
		RealParserFunc: p,
	}
}

type NilParser struct{}

func (n *NilParser) Parse(contents []byte, url string) ParserResult {
	return ParserResult{}

}

func (n *NilParser) Serialize() (string, interface{}) {
	return "NilParser", nil
}
