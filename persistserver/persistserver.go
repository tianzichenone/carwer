package main

import (
	"gopkg.in/olivere/elastic.v5"
	"carwer/rpccommon/rpcsupport"
	"carwer/config"
	"log"
	"carwer/persistserver/rpc"
)

func main() {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config.ESAddress))
	if err != nil {
		panic(err)
	}
	service := persistserver.PersistService{
		Client: client,
	}
	log.Fatal(rpcsupport.CreateRpcServer(&service, config.SaverAddress))

}