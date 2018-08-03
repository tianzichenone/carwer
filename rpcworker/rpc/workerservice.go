package rpcworker

import (
	"carwer/enginee"
	"carwer/fetcher"
	"fmt"
)

type WorkerService struct {
	FunMap map[string]enginee.ParserFunc
}

type WorkParams struct {
	URL        string
	ParserFunc string
}

func (w *WorkerService) Work(params WorkParams, result *enginee.ParserResult) error {
	contents, err := fetcher.Fetcher(params.URL)
	//log.Printf("Fetching Url: %s", req.URL)
	if err != nil {
		//log.Printf("Fetch URL: %s error: %v", req.URL, err)
		return err
	}
	ParserFunc, ok := w.FunMap[params.ParserFunc]
	if !ok {
		return fmt.Errorf("Could not Found ParserFunc for :%s", ParserFunc)
	}
	parserResults := ParserFunc(contents, params.URL)
	*result = parserResults
	return nil
}
