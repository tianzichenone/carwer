package main

import (
	"carwer/rpcworker/rpc"
	"carwer/enginee"
	"carwer/zhenai/parser"
	"carwer/rpccommon/rpcsupport"
	"carwer/config"
	"log"
)

func main() {
	service := rpcworker.WorkerService{
		FunMap: map[string] enginee.ParserFunc{
			"ParserCityList" : parser.ParserCityList,
			"ParserCity": parser.ParserCity,
			//"ParserProfile": parser.ParserProfile,
		},
	}
	log.Fatal(rpcsupport.CreateRpcServer(&service, config.WorkerAddress))
}
