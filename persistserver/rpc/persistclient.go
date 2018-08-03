package persistserver

import (
	"carwer/model"
	"log"
	"carwer/rpccommon/rpcsupport"
	"carwer/config"
)

func MakeItermSaver() chan model.Item {
	ItemChan := make(chan model.Item)
	client, err := rpcsupport.CreateRpcClient(config.SaverAddress)
	if err != nil {
		panic(err)
	}
	result := 0
	go func() {
		itemCount := 0
		for {
			item := <-ItemChan
			err := client.Call(config.ItemSaverRPC, item, &result)
			if err != nil {
				log.Printf("Call rpc error, item: %v", item)
				continue
			}
			if result == 1 {
				itemCount = itemCount + 1
			}

			log.Printf("Getting Items: %v, item count :%d", item, itemCount)
		}
	}()
	return ItemChan
}