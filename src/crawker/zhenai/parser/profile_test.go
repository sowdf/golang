package parser

import (
	"crawker/engine"
	"io/ioutil"
	"model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	//contents, e := fetcher.Fetch("http://album.zhenai.com/u/1847904784")
	//fmt.Printf("%s\n",contents)
	contents, e := ioutil.ReadFile("profile_test_data.html")
	if e != nil {
		panic(e)
	}
	createFile(contents, "profile_test_data.html")
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

	parseResult := ParseProfile(contents, "征婚", "http://album.zhenai.com/u/1847904784")

	if len(parseResult.Items) != 1 {
		t.Errorf("Items should contain 1 element ; but was %v", parseResult.Items)
	}

	actual := parseResult.Items[0]

	if actual != expected {
		t.Errorf("expected is %v:but was %v;", expected, actual)
	}
}
