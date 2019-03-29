package parser

import (
	"encoding/json"
	"model"
	"qupu/common"
	"qupu/fetcher"
	"testing"
)

const filename = "song_test_data.html"

func TestParseSong(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.qupu123.com/minge/sanzi/p316606.html")

	if err != nil {
		panic(err)
	}

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

	common.CreateFile(filename, contents)

	parseResult := ParseSong(contents, "http://www.qupu123.com/tongsu/sanzi/p336145.html")
	expectedByte, err := json.Marshal(expected)
	expectedStr := string(expectedByte)
	if err != nil {
		panic(err)
	}
	actual, err := json.Marshal(parseResult.Item)
	actualStr := string(actual)
	if err != nil {
		panic(err)
	}

	if expectedStr != actualStr {
		t.Errorf("取到的是%v\n,实际的是%v", actualStr, expectedStr)
	}
}
