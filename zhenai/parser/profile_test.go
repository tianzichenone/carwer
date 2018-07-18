package parser

import (
	"testing"
	"io/ioutil"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data")
	if err != nil {
		panic(err)
	}
	ParserProfile(contents, "")
}
