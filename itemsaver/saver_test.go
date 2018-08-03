package itemsaver

import (
	"testing"
	"carwer/model"
	"gopkg.in/olivere/elastic.v5"
	"golang.org/x/net/context"
	"encoding/json"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
		Name: "暮雨而桐",
		Age: 27,
		Height: "172CM",
		Income: "3001-5000元",
		Status: "未婚",
		Edu: "大学本科",
		Occupation: "小学教师",
		Constellation: "魔羯座",
		Hukou: "陕西安康",
		Place: "陕西安康",
	}
	saveItem := model.Item{
		URL: "http://album.zhenai.com/u/1928089545",
		ID: "1928089545",
		TYPE: "zhenai",
		Playload: profile,
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://192.168.134.59:9200"))
	if err != nil {
		panic(err)
	}

	save(saveItem, client)
	resp, err := client.Get().
		Index("data_profile").
		Type("zhenai").
		Id("1928089545").Do(context.Background())
	if err != nil {
		panic(err)
	}
	actualItem := model.Item{}
	err = json.Unmarshal(*resp.Source, &actualItem)
	if err != nil {
		panic(err)
	}
	t.Logf("Item: %+v", actualItem)
	if saveItem.ID != actualItem.ID {
		t.Errorf("save ID: %s, actual ID :%s", saveItem.ID, actualItem.ID)
	}
	if saveItem.URL != actualItem.URL {
		t.Errorf("save URL: %s, actual URL :%s", saveItem.URL, actualItem.URL)
	}
	actualProfile, err := model.MarshJson(actualItem.Playload)
	if err != nil {
		panic(err)
	}
	if saveItem.Playload != actualProfile {
		t.Errorf("save profile: +%v, actual profile :%+v",
			saveItem.Playload, actualProfile)
	}
}
