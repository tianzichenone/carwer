package main

import (
	"carwer/enginee"
	"carwer/persister/client"
	"carwer/rpcsupport"
	"carwer/scheduler"
	"carwer/worker/client"
	"carwer/zhenai/parser"
	"flag"
	"net/rpc"
	"strings"
	"github.com/cloudflare/cfssl/log"
)

var (
	workerHosts = flag.String("worker hosts", "", "The workers hosts,split by comma")
	persisterHost = flag.String("persister host", "", "The persister hosts")
)

func main() {
	flag.Parse()
	if *workerHosts == "" || *persisterHost == "" {
		log.Error("The worker or persister host should set")
		return
	}
	pool := createClientPool(strings.Split(*workerHosts, ","))
	wokerProcess := client.CreateCrawerProcess(pool)

	concurrency := enginee.Concurrency{
		Scheduler:     &scheduler.QueuedScheduler{},
		WorkCount:     100,
		ItemChan:      persister.MakeItermSaver(*persisterHost),
		WorkerProcess: wokerProcess,
	}
	concurrency.Run(enginee.Request{
		URL:    "http://www.zhenai.com/zhenghun",
		Parser: enginee.NewParserFuncFactory("ParserCityList", parser.ParserCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts {
		c, err := rpcsupport.CreateRpcClient(host)
		if err == nil {
			clients = append(clients, c)
		}
	}
	clientChan := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				clientChan <- c
			}
		}

	}()
	return clientChan
}
