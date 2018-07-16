package parser

import (
	"regexp"
	"carwer/enginee"
	"fmt"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([^<]+)</td>`)

func ParserProfile(b []byte) enginee.ParserResult {
	parserResult := enginee.ParserResult{}
	ageMatch := ageRe.FindSubmatch(b)
	fmt.Printf("Getting age:%s\n", ageMatch[1])
	return parserResult
}
