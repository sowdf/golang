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
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1847904784",
		Type: "zhenai",
		Id:   "1847904784",
		Payload: model.Profile{
			Name:       "征婚",
			Gender:     "男士",
			Age:        47,
			Height:     175,
			Weight:     77,
			Income:     "2-5万",
			Marriage:   "离异",
			Education:  "大学本科",
			Occupation: "不能获取",
			Hokou:      "福建厦门",
			Xingzuo:    "射手座(11.22-12.21)",
			Car:        "已买车",
			House:      "已购房",
		},
	}

	// Save expected item
	err := save(expected)

	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	result, err := client.Get().
		Index("data_profile").
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
