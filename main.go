package main

import (
	"net/http"
	"fmt"
	"regexp"
	"carwer/fetcher"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Getting wrong http code: %d", resp.StatusCode)
		return
	}
	body, err := fetcher.Fetcher(resp.Body)
	if err != nil {
		panic(err)
	}
	parserListCity(body)

}


func parserListCity(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	for _, match := range re.FindAllSubmatch(contents, -1) {
		fmt.Printf("City:%s, Url:%s\n", match[2], match[1])
	}
}
