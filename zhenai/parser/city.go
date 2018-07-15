package parser

import (
	"carwer/enginee"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://[a-z]+.zhenai.com/u/[0-9])+"[^>]*>([^<])+</a></th>`)

func ParserCity(b []byte) enginee.ParserResult {
	parserResult := enginee.ParserResult{}
	for _, match := range cityRe.FindSubmatch(b) {
		parserResult.Items = append(parserResult.Items, string(match[2]))
		parserResult.Requests = append(parserResult.Requests, enginee.Request{
			URL: string(match[1]),
			ParserFunc: enginee.NilParserFunc,
		})
	}
	return parserResult

}