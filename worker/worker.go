package main

import (
	"carwer/rpcsupport"
	"carwer/worker/server"
	"flag"
	"log"
)

var (
	workerAddress = flag.String("address", "", "worker server address")
)

func main() {
	flag.Parse()
	if *workerAddress == "" {
		log.Fatal("Worker server address should not None!")
	}
	log.Fatal(rpcsupport.CreateRpcServer(&server.WorkerService{}, *workerAddress))
}
