package parser

import (
	"testing"
	"carwer/fetcher"
)

func TestParserProfile(t *testing.T) {
	contents, err := fetcher.Fetcher("http://album.zhenai.com/u/1558719774")
	if err != nil {
		panic(err)
	}
	ParserProfile(contents)

}
