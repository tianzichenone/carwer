package parser

import (
	"io/ioutil"
	"testing"
)

func TestCityListParser(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data")
	if err != nil {
		panic(err)
	}
	expected_city := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	expected_url := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	result := ParserCityList(contents)
	for i, city := range expected_city {
		if city != result.Items[i] {
			t.Errorf(
				"Expected city :%s,"+
					" but atctual city : %s",
				city, result.Items[i])
		}
	}
	for i, url := range expected_url {
		if url != result.Requests[i].URL {
			t.Errorf("Expected URL: %s, but actual URL: %s",
				url, result.Requests[i].URL)
		}
	}

}
