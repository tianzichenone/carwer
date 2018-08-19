package server

import (
	"carwer/config"
	"carwer/enginee"
	"carwer/model"
	"carwer/zhenai/parser"
	"errors"
	"fmt"
	"log"
)

type SerializeParser struct {
	Name string
	Args interface{}
}

type Request struct {
	URL    string
	Parser SerializeParser
}

type ParserResult struct {
	Items    []model.Item
	Requests []Request
}

func SerializeRequest(req enginee.Request) Request {
	funcName, args := req.Parser.Serialize()
	return Request{
		URL: req.URL,
		Parser: SerializeParser{
			Name: funcName,
			Args: args,
		},
	}
}
func SerializeParserResult(result enginee.ParserResult) ParserResult {
	parserResult := ParserResult{
		Items: result.Items,
	}
	for _, req := range result.Requests {
		parserResult.Requests = append(parserResult.Requests,
			SerializeRequest(req))
	}
	return parserResult
}

func DeserializeRequest(req Request) (enginee.Request, error) {
	engineeParser, err := deserializeParser(req.Parser)
	if err != nil {
		return enginee.Request{}, err
	}
	return enginee.Request{
		URL:    req.URL,
		Parser: engineeParser,
	}, nil
}

func DeserializeParserResult(result ParserResult) enginee.ParserResult {
	parserResult := enginee.ParserResult{
		Items: result.Items,
	}
	for _, req := range result.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("DeserializeRequest faild, req: %v, err:%v", req, err)
			continue
		}
		parserResult.Requests = append(parserResult.Requests,
			request)
	}
	return parserResult
}

func deserializeParser(s SerializeParser) (enginee.Parser, error) {
	switch s.Name {
	case config.ParserCityList:
		return enginee.NewParserFuncFactory(config.ParserCityList,
			parser.ParserCityList), nil
	case config.ParserCity:
		return enginee.NewParserFuncFactory(config.ParserCity,
			parser.ParserCity), nil
	case config.NilParser:
		return &enginee.NilParser{}, nil
	case config.ProfileParser:
		if name, ok := s.Args.(string); !ok {
			return nil, fmt.Errorf("Change args: %v to string error", s.Args)
		} else {
			return parser.NewProfileParser(name), nil
		}
	default:
		return nil, errors.New("No match parser func")

	}
}
