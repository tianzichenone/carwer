package server

import (
	"carwer/config"
	"carwer/rpcsupport"
	"fmt"
	"testing"
	"time"
)

func TestRpcWorker(t *testing.T) {
	go rpcsupport.CreateRpcServer(&WorkerService{},
		config.WorkerAddress)
	time.Sleep(1)
	client, err := rpcsupport.CreateRpcClient(config.WorkerAddress)
	if err != nil {
		panic(err)
	}
	request := Request{
		URL: "http://album.zhenai.com/u/1928089545",
		Parser: SerializeParser{
			Name: config.ParserProfile,
			Args: "暮雨而桐",
		},
	}
	var result ParserResult
	err = client.Call("WorkerService.Process", request, &result)
	realParserResult := DeserializeParserResult(request)
	fmt.Println(realParserResult, err)

}
