package parser

import (
	"io/ioutil"
	"model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	//contents, e := fetcher.Fetch("http://album.zhenai.com/u/1157402933")
	//fmt.Printf("%s\n",contents)
	contents, e := ioutil.ReadFile("profile_test_data.html")
	if e != nil {
		panic(e)
	}
	//createFile(contents,"profile_test_data.html")
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

	parseResult := ParseProfile(contents, "小萍")

	profile := parseResult.Items[0].(model.Profile)

	if profile != expected {
		t.Errorf("expected is %v:but was %v;", expected, profile)
	}
}
