package parser

import (
	"testing"
	"carwer/fetcher"
)

func TestParserProfile(t *testing.T) {
	contents, err := fetcher.Fetcher("http://album.zhenai.com/u/1678110663")
	if err != nil {
		panic(err)
	}
	ParserProfile(contents)

}
