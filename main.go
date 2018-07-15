package main

import (
	"carwer/enginee"
	"carwer/zhenai/parser"
)

func main() {
	enginee.Run(enginee.Request{
		URL: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

}
