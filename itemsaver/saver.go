package itemsaver

import (
	"carwer/model"
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func MakeItermSaver() chan model.Item {
	esClient, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://192.168.134.59:9200"))
	if err != nil {
		panic(err)
	}
	ItemChan := make(chan model.Item)
	go func() {
		itemCount := 0
		for {
			item := <-ItemChan
			save(item, esClient)
			itemCount = itemCount + 1
			log.Printf("Getting Items: %v, item count :%d", item, itemCount)
		}
	}()
	return ItemChan
}

// when we do call in local
func save(item model.Item, client *elastic.Client) {
	resp, err := client.Index().
		Index("data_profile").
		Type(item.TYPE).
		Id(item.ID).
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
