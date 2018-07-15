package fetcher

import (
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func Fetcher(r io.Reader) ([]byte, error) {

	utfEcoding := determineEncoding(r)
	reader := transform.NewReader(r, utfEcoding.NewDecoder())
	return ioutil.ReadAll(reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	contents, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(contents, "")
	return e
}

