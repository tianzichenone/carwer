package enginee

import (
	"carwer/fetcher"
	"log"
)

type Simple struct {

}
func (s Simple) Run(seds ...Request) {
	requests := []Request{}
	for _, request := range seds {
		requests = append(requests, request)
	}
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]
		parserResult, err := Work(req)
		if err != nil {
			continue
		}
		for _, item := range parserResult.Items {
			log.Printf("Get item: %v", item)
		}
		requests = append(requests, parserResult.Requests ...)
	}
}

func Work(req Request) (ParserResult, error) {
	contents, err := fetcher.Fetcher(req.URL)
	//log.Printf("Fetching Url: %s", req.URL)
	if err != nil {
		//log.Printf("Fetch URL: %s error: %v", req.URL, err)
		return ParserResult{}, err
	}
	parserResults := req.ParserFunc(contents)
	return  parserResults, nil
}

