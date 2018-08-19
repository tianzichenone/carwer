package main

import (
	"carwer/config"
	"carwer/persister/server"
	"carwer/rpcsupport"
	"flag"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var (
	persisAddress = flag.String("address", "", "The persister server address")
)

func main() {
	flag.Parse()
	if *persisAddress == "" {
		log.Fatal("Persister server address should not None!")
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config.ESAddress))
	if err != nil {
		panic(err)
	}
	service := persister.PersistService{
		Client: client,
	}
	log.Fatal(rpcsupport.CreateRpcServer(&service, *persisAddress))

}
