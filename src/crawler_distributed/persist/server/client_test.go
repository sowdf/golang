package main

import (
	"crawler_distributed/rpcsupport"
	"model"
	"testing"
)

func TestSaveItem(t *testing.T) {
	// start server
	go serviceRpc(":1234", "test1")

	//time.Sleep(time.Second)
	// start client

	client, err := rpcsupport.NewClient(":1234")

	if err != nil {
		panic(err)
	}

	item := model.Spectrum{
		Id:          "336145",
		Type:        "qupu123",
		Title:       "​30女儿情（双谱）",
		From:        "新编中国声乐作品第15集  ",
		FromUrl:     "http://www.qupu123.com/tongsu/sanzi/p336145.html",
		Composition: "许镜清",
		Uploader:    "瓦莲金娜AA  ",
		Views:       3561,
		UploadDate:  "2018-05-21",
		ImageUrls:   []string{"http://www.qupu123.com/Public/Uploads/2018/05/21/5b02162abcb9b.jpg", "http://www.qupu123.com/Public/Uploads/2018/05/21/5b0216332b8d1.jpg"},
		ClickNum:    0,
		Lyrics:      "杨洁",
	}

	var result string

	err = client.Call("ItemSaveServerService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result %s : %v", result, err)
	}

}
