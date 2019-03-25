package persist

import (
	"encoding/json"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name:       "小萍",
		Gender:     "女士",
		Age:        43,
		Height:     166,
		Weight:     0,
		Income:     "3千以下",
		Marriage:   "离异",
		Education:  "高中及以下",
		Occupation: "不能获取",
		Hokou:      "内蒙古阿拉善盟",
		Xingzuo:    "射手座(11.22-12.21)",
		Car:        "",
		House:      "",
	}

	id, err := save(expected)

	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	result, err := client.Get().Index("data_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual model.Profile

	err = json.Unmarshal(*result.Source, &actual)

	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("Got : %v,expexted : %v\n", actual, expected)
	}

}
