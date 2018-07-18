package main

import (
	"carwer/zhenai/parser"
	"carwer/enginee"
	"carwer/scheduler"
)

func main() {
	enginee.Concurrency{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkCount: 100,

	}.Run(enginee.Request{
		URL: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
