package parser

import (
	"carwer/fetcher"
	"fmt"
	"testing"
)

func TestCityParser(t *testing.T) {
	contents, err := fetcher.Fetcher("http://www.zhenai.com/zhenghun/alashanmeng")
	if err != nil {
		panic(err)
	}
	parserResult := ParserCity(contents)
	for _, name := range parserResult.Items {
		fmt.Printf("%s\n", name)
	}
	for _, request := range parserResult.Requests {
		fmt.Printf("%s\n", request.URL)
	}
}
