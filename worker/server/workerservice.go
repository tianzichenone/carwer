package server

import (
	"carwer/enginee"
)

type WorkerService struct {
}

func (w *WorkerService) Process(req Request, result *ParserResult) error {
	request, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	parserResult, err := enginee.Work(request)
	if err != nil {
		return err
	}
	*result = SerializeParserResult(parserResult)
	return nil
}
