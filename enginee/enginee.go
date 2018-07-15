package enginee

import (
	"carwer/fetcher"
	"log"
)

func Run(seds ...Request) {
	requests := []Request{}
	for _, request := range seds {
		requests = append(requests, request)
	}
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]
		contents, err := fetcher.Fetcher(req.URL)
		log.Printf("Fetching Url: %s", req.URL)
		if err != nil {
			log.Printf("Fetch URL: %s error", req.URL)
		}
		parserResults := req.ParserFunc(contents)
		for _, item := range parserResults.Items {
			log.Printf("Get item: %v", item)
		}
		requests = append(requests, parserResults.Requests ...)
	}
}
