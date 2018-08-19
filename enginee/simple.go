package enginee

import (
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
		requests = append(requests, parserResult.Requests...)
	}
}
