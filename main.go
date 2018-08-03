package main

import (
	"carwer/zhenai/parser"
	"carwer/enginee"
	"carwer/scheduler"
	"carwer/persistserver/rpc"
)

func main() {
	enginee.Concurrency{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 100,
		ItemChan: persistserver.MakeItermSaver(),

	}.Run(enginee.Request{
		URL: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
