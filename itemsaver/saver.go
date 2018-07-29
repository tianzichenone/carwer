package itemsaver

import (
	"log"
	"carwer/model"
)

func MakeItermSaver() chan model.Item {
	ItemChan := make(chan model.Item)
	go func() {
		itemCount := 0
		for {
			item := <-ItemChan
			//Save(item)
			itemCount = itemCount + 1
			log.Printf("Getting Items: %v, item count :%d", item, itemCount)
		}
	}()
	return ItemChan
}

/*
func Save(item interface{}) {
	esClient, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := esClient.Index().
		Index("data_profile").
		Type("zhenai").
		Id().
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}*/