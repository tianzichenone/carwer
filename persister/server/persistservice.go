package persister

import (
	"carwer/model"
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
)

type PersistService struct {
	Client *elastic.Client
}

func (ps *PersistService) Save(item model.Item, ok *int) error {
	resp, err := ps.Client.Index().
		Index("data_profile").
		Type(item.TYPE).
		Id(item.ID).
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		*ok = 0
		return err
	}
	fmt.Println(resp)
	*ok = 1
	return nil

}
