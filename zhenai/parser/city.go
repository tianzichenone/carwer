package parser

import (
	"carwer/enginee"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://[a-z]+.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	nextRe = regexp.MustCompile(`href="([^>]+")>下一页`)
	otherRe = regexp.MustCompile(`<a target="_blank" href="(http://www.zhenai.com/zhenghun/[^>]+)">([^<]+)</a>`)
)

func ParserCity(b []byte) enginee.ParserResult {
	parserResult := enginee.ParserResult{}
	for _, match := range profileRe.FindAllSubmatch(b, -1) {
		name := string(match[2])
		url  := string(match[1])
		//parserResult.Items = append(parserResult.Items, string(match[2]))
		parserResult.Requests = append(parserResult.Requests, enginee.Request{
			URL: string(match[1]),
			ParserFunc: func(contents []byte) enginee.ParserResult {
				return ParserProfile(contents, name, url)
			},
		})
	}
	for _, match := range nextRe.FindAllSubmatch(b, -1) {
		//parserResult.Items = append(parserResult.Items, "下一页")
		parserResult.Requests = append(parserResult.Requests, enginee.Request{
			URL: string(match[1]),
			ParserFunc: ParserCity,
		})
	}
	for _, match := range otherRe.FindAllSubmatch(b, -1) {
		//parserResult.Items = append(parserResult.Items, string(match[2]))
		parserResult.Requests = append(parserResult.Requests, enginee.Request{
			URL: string(match[1]),
			ParserFunc: ParserCity,
		})
	}
	parserResult.Items = append(parserResult.Items)
	return parserResult

}
