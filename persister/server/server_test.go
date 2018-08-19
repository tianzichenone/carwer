package persister

import (
	"carwer/config"
	"carwer/model"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"testing"
)

func TestPersisService_Save(t *testing.T) {
	conn, err := net.Dial("tcp", ":1235")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result int
	profile := model.Profile{
		Name:          "暮雨而桐",
		Age:           27,
		Height:        "172CM",
		Income:        "3001-5000元",
		Status:        "未婚",
		Edu:           "大学本科",
		Occupation:    "小学教师",
		Constellation: "魔羯座",
		Hukou:         "陕西安康",
		Place:         "陕西安康",
	}
	saveItem := model.Item{
		URL:      "http://album.zhenai.com/u/1928089545",
		ID:       "1928089545",
		TYPE:     "zhenai",
		Playload: profile,
	}
	err = client.Call(config.ItemSaverRPC, saveItem, &result)
	fmt.Println(result, err)
}
