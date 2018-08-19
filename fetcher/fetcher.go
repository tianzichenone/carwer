package fetcher

import (
	"bufio"
	"fmt"
	"github.com/emicklei/go-restful/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func Fetcher(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Getting wrong http code: %d", resp.StatusCode)
	}
	newReader := bufio.NewReader(resp.Body)
	utfEcoding := determineEncoding(newReader)
	utf8Reader := transform.NewReader(newReader, utfEcoding.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	contents, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(contents, "")
	return e
}
