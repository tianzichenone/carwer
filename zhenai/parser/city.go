package parser

import (
	"carwer/enginee"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://[a-z]+.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	nextRe    = regexp.MustCompile(`href="([^>]+")>下一页`)
	otherRe   = regexp.MustCompile(`<a target="_blank" href="(http://www.zhenai.com/zhenghun/[^>]+)">([^<]+)</a>`)
)

func ParserCity(b []byte, _ string) enginee.ParserResult {
	parserResult := enginee.ParserResult{}
	for _, match := range profileRe.FindAllSubmatch(b, -1) {
		//parserResult.Items = append(parserResult.Items, string(match[2]))
		parserResult.Requests = append(parserResult.Requests, enginee.Request{
			URL:    string(match[1]),
			Parser: NewProfileParser(string(match[2])),
		})
	}
	for _, match := range nextRe.FindAllSubmatch(b, -1) {
		//parserResult.Items = append(parserResult.Items, "下一页")
		parserResult.Requests = append(parserResult.Requests, enginee.Request{
			URL:    string(match[1]),
			Parser: enginee.NewParserFuncFactory("ParserCity", ParserCity),
		})
	}
	for _, match := range otherRe.FindAllSubmatch(b, -1) {
		//parserResult.Items = append(parserResult.Items, string(match[2]))
		parserResult.Requests = append(parserResult.Requests, enginee.Request{
			URL:    string(match[1]),
			Parser: enginee.NewParserFuncFactory("ParserCity", ParserCity),
		})
	}
	//parserResult.Items = append(parserResult.Items)
	return parserResult

}

/*
func CreateParerFunc(name string) func([]byte, string) enginee.ParserResult {
	return func(contents []byte, url string) enginee.ParserResult {
		return ParserProfile(contents, name, url)
	}
}*/

type ProfileParser struct {
	Username string
}

func (p *ProfileParser) Parse(contents []byte, url string) enginee.ParserResult {
	return parserProfile(contents, p.Username, url)

}

func (p *ProfileParser) Serialize() (string, interface{}) {
	return "ProfileParser", p.Username
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		Username: name,
	}
}
