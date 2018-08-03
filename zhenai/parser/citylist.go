package parser

import (
	"carwer/enginee"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)

func ParserCityList(b []byte, _ string) enginee.ParserResult {
	matches := cityListRe.FindAllSubmatch(b, -1)
	result := enginee.ParserResult{}
	for _, match := range matches {
		//result.Items = append(result.Items, string(match[2]))
		result.Requests = append(result.Requests, enginee.Request{
			URL: string(match[1]),
			ParserFunc: ParserCity,
		})
	}
	//result.Items = append(result.Items)
	return result

}