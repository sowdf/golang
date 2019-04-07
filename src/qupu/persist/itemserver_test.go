package persist

import (
	"crawker/engine"
	"encoding/json"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := model.Spectrum{
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
	const index = "dating_test"
	// Save expected item

	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	err = Save(client, expected, index)

	if err != nil {
		panic(err)
	}

	result, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item

	err = json.Unmarshal(*result.Source, &actual)

	profile, _ := model.FromJsonObj(actual.Payload)

	actual.Payload = profile

	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("Got : %v,expexted : %v\n", actual, expected)
	}

}
