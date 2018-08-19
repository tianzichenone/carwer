package client

import (
	"carwer/config"
	"carwer/enginee"
	"carwer/worker/server"
	"github.com/emicklei/go-restful/log"
	"net/rpc"
)

func CreateCrawerProcess(client chan *rpc.Client) enginee.CrawerProcess {

	return func(request enginee.Request) (enginee.ParserResult, error) {
		req := server.SerializeRequest(request)
		var result server.ParserResult
		c := <-client
		err := c.Call(config.WorkerRPC, req, &result)
		if err != nil {
			log.Printf("Call worker req: %v rpc error:%v", request, err)
			return enginee.ParserResult{}, err
		}

		return server.DeserializeParserResult(result), nil

	}
}
