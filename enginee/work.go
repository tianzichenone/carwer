package enginee

import "carwer/fetcher"

func Work(req Request) (ParserResult, error) {
	contents, err := fetcher.Fetcher(req.URL)
	//log.Printf("Fetching Url: %s", req.URL)
	if err != nil {
		//log.Printf("Fetch URL: %s error: %v", req.URL, err)
		return ParserResult{}, err
	}
	parserResults := req.ParserFunc(contents, req.URL)
	return  parserResults, nil
}
